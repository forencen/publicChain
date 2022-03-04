package main

import (
	"fmt"
	"publicChain/block"
	"publicChain/pow"
)

func main() {
	bl := block.CreateGenesisBlock("test")
	bl.SetMinerInfo(pow.NewProofOfWork(bl).Run())
	fmt.Printf("%x\n", bl.Hash)
	fmt.Printf("%d\n", bl.Nonce)
}
