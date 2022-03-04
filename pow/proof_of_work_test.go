package pow

import (
	"fmt"
	"publicChain/block"
	"testing"
)

func TestNewBlock(t *testing.T) {
	genesisBlock := block.CreateGenesisBlock("Genenis Block")
	genesisBlock.SetMinerInfo(NewProofOfWork(genesisBlock).Run())
	fmt.Printf("%x\n", genesisBlock.Hash)
	fmt.Printf("%d\n", genesisBlock.Nonce)

	sBytes := genesisBlock.Serialize()
	fmt.Println(sBytes)
	block1 := block.Deserialize(sBytes)
	fmt.Printf("%x\n", block1.Hash)
}
