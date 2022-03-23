package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"publicChain/block"
	"publicChain/pow"
	"publicChain/transaction"
	"publicChain/utils"
)

type Cli struct {
}

func (c *Cli) printUsage() {
	fmt.Println("this is usage")
}

// 创建创世区块
func (c *Cli) genesisBlock2Db(address string) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	txs := []*transaction.Transaction{transaction.NewCoinbaseTransaction(address)}
	genesisBlock := block.CreateGenesisBlock(txs)
	pow.NewProofOfWork(genesisBlock).Run()
	bc.AddBlockInstanceToBlockChan(genesisBlock)
}

func send(from []string, to []string, amount []string) {

}

func (c *Cli) addBlock(txs []*transaction.Transaction) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	if bc.Tip == nil {
		return
	}
	nowBlockBytes, _ := bc.Db.Get(bc.Tip)
	nowBlock := block.Deserialize(nowBlockBytes)
	newBlock := block.NewBlock(nowBlock.Height+1, txs, nowBlock.Hash)
	pow.NewProofOfWork(newBlock).Run()
	bc.AddBlockToBlockChan(newBlock)
}

func (c *Cli) printChain() {
	bc := block.BlockChainObject()
	bc.PrintChain()
	defer bc.Db.Close()
}

func (c *Cli) Run() {

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockChainCmd := flag.NewFlagSet("createBlockChainCmd", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "xxxxxx", "区块内容")
	flagCreateBlockChaiData := createBlockChainCmd.String("address", "", "创世区块创建者的地址")

	flagFrom := sendCmd.String("from", "", "给钱")
	flagTo := sendCmd.String("to", "", "收钱")
	flagAmount := sendCmd.String("amount", "", "钱")

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
	case "send":
		err := sendCmd.Parse(os.Args[2:])
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
		fmt.Println(*flagAddBlockData)
		c.addBlock(nil)
	}
	if printChainCmd.Parsed() {
		c.printChain()
	}
	if createBlockChainCmd.Parsed() {
		c.genesisBlock2Db(*flagCreateBlockChaiData)
	}
	if sendCmd.Parsed() {
		fmt.Println(utils.Json2StrArray(*flagFrom))
		fmt.Println(utils.Json2StrArray(*flagTo))
		fmt.Println(utils.Json2StrArray(*flagAmount))
	}
}
