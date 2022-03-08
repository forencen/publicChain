package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

func AddBlock2Db(bc *block.BlockChain, data string) []byte {
	nowBlockBytes, _ := bc.Db.Get(bc.Tip)
	nowBlock := block.Deserialize(nowBlockBytes)
	newBlock := block.NewBlock(nowBlock.Height+1, data, nowBlock.Hash)
	pow.NewProofOfWork(newBlock).Run()
	bc.AddBlockToBlockChan(newBlock)
	return newBlock.Hash
}

func main() {
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "xxxxxx", "区块内容")
	if len(os.Args) <= 1 {
		printUsage()
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
	default:
		printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		fmt.Println(*flagAddBlockData)
	}

	//bc := block.NewBlockChain()
	//defer bc.Db.Close()

	//genesisBlock2Db(bc)
	//AddBlock2Db(bc, "a to b 100")
	//AddBlock2Db(bc, "b to a 30")
	//AddBlock2Db(bc, "c to b 100")
	//AddBlock2Db(bc, "b to d 10")
	//bc.PrintChain()
}

func printUsage() {
	fmt.Println("this is usage")
}
