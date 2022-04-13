package cli

import (
	"fmt"
	"publicChain/block"
)

func (c *Cli) getBalance(address string) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	amount := bc.GetBalance(address)
	fmt.Printf("%s balance: %d\n", address, amount)
}
