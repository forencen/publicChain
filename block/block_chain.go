package block

import (
	"fmt"
	"log"
	"math/big"
	"publicChain/db"
	"time"
)

type BlockChain struct {
	Tip []byte //最新的区块hash
	Db  *db.DbHelper
}

const dbName = "block.db"

type genBlockFun func(bc *BlockChain) *Block

func NewBlockChain(gen genBlockFun) *BlockChain {
	blockDb := db.NewDbHelper(dbName)
	result, err := blockDb.Get([]byte("Tip"))
	bc := &BlockChain{result, blockDb}
	if err != nil {
		bc.AddBlockInstanceToBlockChan(gen(bc))
	}
	return bc
}

func (bc *BlockChain) AddBlockInstanceToBlockChan(genBlock *Block) {
	bc.Tip = genBlock.Hash
	err := bc.Db.Put(bc.Tip, genBlock.Serialize())
	if err != nil {
		log.Panic(err)
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
		fmt.Printf("%x\t", block.Hash)
		fmt.Printf("%d\t", block.Height)
		fmt.Printf("%s\t", string(block.Data))
		fmt.Printf("%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"))
		if big.NewInt(0).Cmp(new(big.Int).SetBytes(block.PrevBlockHash)) == 0 {
			break
		}
	}
}
