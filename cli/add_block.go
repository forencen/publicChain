package cli

import (
	"publicChain/block"
	"publicChain/transaction"
)

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
