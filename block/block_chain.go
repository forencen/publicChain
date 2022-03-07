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

func NewBlockChain() *BlockChain {
	blockDb := db.NewDbHelper(dbName)
	result, err := blockDb.Get([]byte("Tip"))
	if err != nil {
		return nil
	}
	return &BlockChain{result, blockDb}
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

func (bc *BlockChain) PrintChain(lastHash []byte) {
	lastBlockBytes, _ := bc.Db.Get(lastHash)
	lastBlock := Deserialize(lastBlockBytes)
	fmt.Printf("%x\t", lastBlock.Hash)
	fmt.Printf("%d\t", lastBlock.Height)
	fmt.Printf("%s\n", time.Unix(lastBlock.Timestamp, 0).Format("2006-01-02 15:04:05"))
	if big.NewInt(0).Cmp(new(big.Int).SetBytes(lastBlock.PrevBlockHash)) == 0 {
		return
	}
	bc.PrintChain(lastBlock.PrevBlockHash)
}
