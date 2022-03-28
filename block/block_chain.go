package block

import (
	"math/big"
	"publicChain/db"
	"publicChain/pow"
	"publicChain/transaction"
)

type BlockChain struct {
	Tip []byte //最新的区块hash
	Db  *db.DbHelper
}

const dbName = "block.db"

func BlockChainObject() *BlockChain {
	blockDb := db.NewDbHelper(dbName)
	result, _ := blockDb.Get([]byte("Tip"))
	return &BlockChain{result, blockDb}
}

func (bc *BlockChain) LastBlock() *Block {
	if bc.Tip == nil {
		panic("please genesis block")
	}
	get, err := bc.Db.Get(bc.Tip)
	if err != nil {
		panic("db tip is error")
	}
	return Deserialize(get)
}

func (bc *BlockChain) AddBlockInstanceToBlockChan(genBlock *Block) {
	if bc.Tip != nil {
		return
	}
	bc.Tip = genBlock.Hash
	batch := bc.Db.NewBatch()
	batch.Put(bc.Tip, genBlock.Serialize())
	batch.Put([]byte("Tip"), bc.Tip)
	if err := batch.Commit(); err != nil {
		batch.Rollback()
	}
}

func (bc *BlockChain) AddBlockToBlockChan(block *Block) {
	putErr := bc.Db.Put(block.Hash, block.Serialize())
	if putErr != nil {
		return
	}
	bc.Tip = block.Hash
	saveTipErr := bc.Db.Put([]byte("Tip"), bc.Tip)
	if saveTipErr != nil {
		return
	}
}

func (bc *BlockChain) Iterator() *ChainIterator {
	return &ChainIterator{bc.Tip, bc.Db}
}

func (bc *BlockChain) PrintChain() {
	iterator := bc.Iterator()
	var block *Block
	for {
		block = iterator.Next()
		if block == nil {
			return
		}
		block.PrintBlock()
		if big.NewInt(0).Cmp(new(big.Int).SetBytes(block.PrevBlockHash)) == 0 {
			break
		}
	}
}

func (bc *BlockChain) GetUnUseUtxo(address string) []*transaction.Utxo {
	iterator := bc.Iterator()
	var block *Block
	var utxoTemp *transaction.Utxo
	utxoMapping := make(map[string]*transaction.Utxo)
	for {
		block = iterator.Next()
		if block == nil {
			break
		}
		for _, tx := range block.Txs {

			for _, vin := range tx.Vins {
				if vin.UnLockWithAddress(address) {
					utxoTemp = transaction.NewUtxo(vin.TxHash, vin.Vout, 0, address, true)
					utxoMapping[utxoTemp.Index()] = utxoTemp
				}
			}
			for i, vout := range tx.Vouts {
				if vout.UnLockWithAddress(address) {
					utxoTemp = transaction.NewUtxo(tx.Hash, i, vout.Value, address, false)
					if mappingItem, ok := utxoMapping[utxoTemp.Index()]; ok {
						mappingItem.Value = vout.Value
					} else {
						utxoMapping[utxoTemp.Index()] = utxoTemp
					}

				}

			}
		}

	}
	utxoList := make([]*transaction.Utxo)
	utxoMapping
	return
}

func (bc *BlockChain) GetBalance(address string) int64 {
	iterator := bc.Iterator()
	var block *Block
	var utxoTemp *transaction.Utxo
	utxoMapping := make(map[string]*transaction.Utxo)
	for {
		block = iterator.Next()
		if block == nil {
			break
		}
		for _, tx := range block.Txs {

			for _, vin := range tx.Vins {
				if vin.UnLockWithAddress(address) {
					utxoTemp = transaction.NewUtxo(vin.TxHash, vin.Vout, 0, address, true)
					utxoMapping[utxoTemp.Index()] = utxoTemp
				}
			}
			for i, vout := range tx.Vouts {
				if vout.UnLockWithAddress(address) {
					utxoTemp = transaction.NewUtxo(tx.Hash, i, vout.Value, address, false)
					if mappingItem, ok := utxoMapping[utxoTemp.Index()]; ok {
						mappingItem.Value = vout.Value
					} else {
						utxoMapping[utxoTemp.Index()] = utxoTemp
					}

				}

			}
		}

	}
	amount := int64(0)
	for _, value := range utxoMapping {
		if !value.IsUsed {
			amount += value.Value
		}
	}
	return amount
}

// MineNewBlock 挖掘新的区块
func (bc *BlockChain) MineNewBlock(from []string, to []string, amount []string) {
	var txs []*transaction.Transaction
	// todo  完成交易对的创建
	lastBlock := bc.LastBlock()
	block := NewBlock(lastBlock.Height+1, txs, lastBlock.Hash)
	p := pow.NewProofOfWork(block)
	p.Run()
	bc.AddBlockToBlockChan(block)
}
