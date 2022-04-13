package block

import (
	"math/big"
	"publicChain/db"
	"publicChain/pow"
	"publicChain/transaction"
	"strconv"
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

// GetUnUseUtxo 获取地址所有的UTXO
func (bc *BlockChain) GetUnUseUtxo(address string) []*transaction.Utxo {
	iterator := bc.Iterator()
	var block *Block
	var usedKeySet = make(map[string]struct{})
	utxoList := make([]*transaction.Utxo, 0, 10)
	for {
		block = iterator.Next()
		if block == nil {
			break
		}
		for i := len(block.Txs) - 1; i >= 0; i-- {
			for _, vin := range block.Txs[i].Vins {
				if vin.UnLockWithAddress(address) {
					usedKeySet[string(vin.TxHash)+strconv.Itoa(vin.Vout)] = struct{}{}
				}
			}

			for index, vout := range block.Txs[i].Vouts {
				if vout.UnLockWithAddress(address) {
					k := string(block.Txs[i].Hash) + strconv.Itoa(index)
					if _, ok := usedKeySet[k]; ok {
						continue
					}
					utxoList = append(utxoList, transaction.NewUtxo(block.Txs[i].Hash, i, vout.Value, address))
				}
			}
		}

	}
	return utxoList
}

func (bc *BlockChain) GetBalance(address string) int64 {
	utxos := bc.GetUnUseUtxo(address)
	var sumAmount int64
	for _, utxo := range utxos {
		sumAmount += utxo.Value
	}
	return sumAmount
}

func (bc *BlockChain) FindAddressEnoughUtxos(address string, amount int64) ([]*transaction.Utxo, bool) {
	iterator := bc.Iterator()
	var block *Block
	var usedKeySet = make(map[string]struct{})
	utxoList := make([]*transaction.Utxo, 0, 10)
	var nowAmount int64
enough:
	for {
		block = iterator.Next()
		if block == nil {
			break
		}
		for i := len(block.Txs) - 1; i >= 0; i-- {
			for _, vin := range block.Txs[i].Vins {
				if vin.UnLockWithAddress(address) {
					usedKeySet[string(vin.TxHash)+strconv.Itoa(vin.Vout)] = struct{}{}
				}
			}

			for index, vout := range block.Txs[i].Vouts {
				if vout.UnLockWithAddress(address) {
					k := string(block.Txs[i].Hash) + strconv.Itoa(index)
					if _, ok := usedKeySet[k]; ok {
						continue
					}
					utxoList = append(utxoList, transaction.NewUtxo(block.Txs[i].Hash, i, vout.Value, address))
					nowAmount += vout.Value
					if nowAmount >= amount {
						break enough
					}
				}
			}
		}

	}
	return utxoList, nowAmount >= amount
}
func (bc *BlockChain) NewSimpleTransaction(from string, to string, amount string) *transaction.Transaction {
	needAmount, _ := strconv.ParseInt(amount, 10, 64)
	needUtxos, isEnough := bc.FindAddressEnoughUtxos(from, needAmount)
	if !isEnough {
		return nil
	}
	var vins []*transaction.TxInput
	var canUseAmount int64
	for _, utxo := range needUtxos {
		canUseAmount += utxo.Value
		vins = append(vins, &transaction.TxInput{TxHash: utxo.TxHash, Vout: utxo.Vout, ScriptSig: utxo.ScriptSig})
	}
	var vouts []*transaction.TxOutput
	vouts = append(vouts, &transaction.TxOutput{ScriptSig: to, Value: needAmount})
	if canUseAmount > needAmount {
		vouts = append(vouts, &transaction.TxOutput{ScriptSig: from, Value: canUseAmount - needAmount})
	}
	tran := &transaction.Transaction{Vins: vins, Vouts: vouts}
	tran.SetTxHash()
	return tran
}

// MineNewBlock 挖掘新的区块
// 多对多的转账情况还未实现
func (bc *BlockChain) MineNewBlock(from []string, to []string, amount []string) {
	var txs []*transaction.Transaction
	if !(len(from) == len(to) && len(to) == len(amount)) {
		return
	}
	for index := range from {
		tempTx := bc.NewSimpleTransaction(from[index], to[index], amount[index])
		if tempTx != nil {
			txs = append(txs, tempTx)
		}
	}
	lastBlock := bc.LastBlock()
	block := NewBlock(lastBlock.Height+1, txs, lastBlock.Hash)
	p := pow.NewProofOfWork(block)
	p.Run()
	bc.AddBlockToBlockChan(block)
}
