package cli

import (
	"fmt"
	"publicChain/wallet"
)

func (c *Cli) createWallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("创建的地址为：%s\n", address)
}
