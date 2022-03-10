package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"publicChain/block"
	"publicChain/pow"
)

type Cli struct {
	BlockChain *block.BlockChain
}

func (c *Cli) printUsage() {
	fmt.Println("this is usage")
}

func (c *Cli) genesisBlock2Db(data string) {
	genesisBlock := block.CreateGenesisBlock(data)
	pow.NewProofOfWork(genesisBlock).Run()
	c.BlockChain.AddBlockInstanceToBlockChan(genesisBlock)
}

func (c *Cli) addBlock(data string) {
	if c.BlockChain.Tip == nil {
		return
	}
	nowBlockBytes, _ := c.BlockChain.Db.Get(c.BlockChain.Tip)
	nowBlock := block.Deserialize(nowBlockBytes)
	newBlock := block.NewBlock(nowBlock.Height+1, data, nowBlock.Hash)
	pow.NewProofOfWork(newBlock).Run()
	c.BlockChain.AddBlockToBlockChan(newBlock)
}

func (c *Cli) printChain() {
	c.BlockChain.PrintChain()
}

func (c *Cli) Run() {

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChainCmd", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "xxxxxx", "区块内容")
	flagCreateBlockChaiData := createBlockChainCmd.String("data", "Genesis data....", "创世区块")
	if len(os.Args) <= 1 {
		c.printUsage()
		return
	}
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createBlockChain":
		if err := createBlockChainCmd.Parse(os.Args[2:]); err != nil {
			log.Panic(err)
		}
	default:
		c.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		c.addBlock(*flagAddBlockData)
	}
	if printChainCmd.Parsed() {
		c.printChain()
	}
	if createBlockChainCmd.Parsed() {
		c.genesisBlock2Db(*flagCreateBlockChaiData)
	}
}
