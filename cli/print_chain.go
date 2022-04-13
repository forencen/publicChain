package cli

import "publicChain/block"

func (c *Cli) printChain() {
	bc := block.BlockChainObject()
	bc.PrintChain()
	defer bc.Db.Close()
}
