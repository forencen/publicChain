package main

import (
	"fmt"
	"publicChain/block"
	"publicChain/pow"
)

func genesisBlock2Db(bc *block.BlockChain) {
	genesisBlock := block.CreateGenesisBlock("test")
	pow.NewProofOfWork(genesisBlock).Run()
	bc.AddBlockInstanceToBlockChan(genesisBlock)
	fmt.Printf("%x\n", genesisBlock.Hash)
	fmt.Printf("%d\n", genesisBlock.Nonce)
}

func AddBlock2Db(bc *block.BlockChain) []byte {
	nowBlockBytes, _ := bc.Db.Get(bc.Tip)
	nowBlock := block.Deserialize(nowBlockBytes)
	newBlock := block.NewBlock(nowBlock.Height+1, "transaction", nowBlock.Hash)
	pow.NewProofOfWork(newBlock).Run()
	bc.AddBlockToBlockChan(newBlock)
	return newBlock.Hash
}

func main() {
	bc := block.NewBlockChain()
	defer bc.Db.Close()

	//genesisBlock2Db(bc)
	AddBlock2Db(bc)
	AddBlock2Db(bc)
	AddBlock2Db(bc)
	AddBlock2Db(bc)
	bc.PrintChain(bc.Tip)
}
