package main

import (
	"publicChain/block"
	"publicChain/cli"
)

func main() {
	bc := block.NewBlockChain()
	c := &cli.Cli{BlockChain: bc}
	c.Run()
}
