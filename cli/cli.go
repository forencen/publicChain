package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"publicChain/block"
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
	bc.AddBlockInstanceToBlockChan(genesisBlock)
}

func (c *Cli) send(from []string, to []string, amount []string) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	bc.MineNewBlock(from, to, amount)
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
	bc.AddBlockToBlockChan(newBlock)
}

func (c *Cli) printChain() {
	bc := block.BlockChainObject()
	bc.PrintChain()
	defer bc.Db.Close()
}

func (c *Cli) getBalance(address string) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	amount := bc.GetBalance(address)
	fmt.Printf("%s balance: %d", address, amount)
}

func (c *Cli) Run() {

	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)

	createBlockChainCmd := flag.NewFlagSet("createBlockChainCmd", flag.ExitOnError)
	flagCreateBlockChaiData := createBlockChainCmd.String("address", "", "创世区块创建者的地址")

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	flagAddBlockData := addBlockCmd.String("data", "xxxxxx", "区块内容")

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	flagFrom := sendCmd.String("from", "", "给钱")
	flagTo := sendCmd.String("to", "", "收钱")
	flagAmount := sendCmd.String("amount", "", "钱")

	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	balanceAddress := getBalanceCmd.String("address", "xxxxxx", "查询地址")

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
	case "getBalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); err != nil {
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
		//fmt.Println(utils.Json2StrArray(*flagFrom))
		//fmt.Println(utils.Json2StrArray(*flagTo))
		//fmt.Println(utils.Json2StrArray(*flagAmount))
		c.send(utils.Json2StrArray(*flagFrom), utils.Json2StrArray(*flagTo), utils.Json2StrArray(*flagAmount))
	}
	if getBalanceCmd.Parsed() {
		c.getBalance(*balanceAddress)
	}
}
