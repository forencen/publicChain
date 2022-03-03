package block

import (
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {
	block := CreateGenesisBlock("Genenis Block")
	fmt.Printf("%x\n", block.Hash)
	fmt.Printf("%d\n", block.Nonce)

	sBytes := block.Serialize()
	fmt.Println(sBytes)
	block1 := Deserialize(sBytes)
	fmt.Printf("%x\n", block1.Hash)
}
