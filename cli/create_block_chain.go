package cli

import (
	"publicChain/block"
	"publicChain/transaction"
)

// 创建创世区块
func (c *Cli) genesisBlock2Db(address string) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	txs := []*transaction.Transaction{transaction.NewCoinbaseTransaction(address)}
	genesisBlock := block.CreateGenesisBlock(txs)
	bc.AddBlockInstanceToBlockChan(genesisBlock)
}
