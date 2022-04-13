package cli

import "publicChain/block"

// send 预计准备多地址发送utxo，多地址收款
// 先实现单地址对单个地址和找零逻辑
func (c *Cli) send(from []string, to []string, amount []string) {
	bc := block.BlockChainObject()
	defer bc.Db.Close()
	bc.MineNewBlock(from[0], to[0], amount[0])
}
