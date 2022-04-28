package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
}

func (c *Cli) printUsage() {
	fmt.Println("this is usage")
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

	createWallet := flag.NewFlagSet("createWallet", flag.ExitOnError)

	if len(os.Args) <= 1 {
		c.printUsage()
		return
	}
	switch os.Args[1] {
	case "createWallet":
		err := createWallet.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
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
	if createWallet.Parsed() {
		c.createWallet()
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
		//c.send(utils.Json2StrArray(*flagFrom), utils.Json2StrArray(*flagTo), utils.Json2StrArray(*flagAmount))
		c.send(*flagFrom, *flagTo, *flagAmount)
	}
	if getBalanceCmd.Parsed() {
		c.getBalance(*balanceAddress)
	}
}
