package main

import (
	"publicChain/block"
	"publicChain/cli"
	"publicChain/pow"
)

func genesisBlock2Db(bc *block.BlockChain) *block.Block {
	genesisBlock := block.CreateGenesisBlock("genesisBlock...")
	pow.NewProofOfWork(genesisBlock).Run()
	bc.AddBlockInstanceToBlockChan(genesisBlock)
	return genesisBlock
}

func main() {
	bc := block.NewBlockChain(genesisBlock2Db)
	c := &cli.Cli{BlockChain: bc}
	c.Run()
}
