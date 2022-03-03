package block

import (
	"fmt"
	"publicChain/pow"
	"testing"
)

func TestNewBlock(t *testing.T) {
	block := CreateGenesisBlock("Genenis Block")
	fmt.Printf("%x\n", block.Hash)
	fmt.Printf("%d\n", block.Nonce)
	pow := pow.NewProofOfWork(block)
	fmt.Println(pow.IsValida())

	sBytes := block.Serialize()
	fmt.Println(sBytes)
	block1 := Deserialize(sBytes)
	fmt.Printf("%x\n", block1.Hash)
}
