package wallet

import (
	"fmt"
	"testing"
)

func TestNewWallet(t *testing.T) {
	w := NewWallet()
	fmt.Printf("%s\n", w.GetAddress())
}

func TestDecode(t *testing.T) {
	address := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
	res := Base58Decode([]byte(address))
	res = append([]byte{version}, res...)
	fmt.Println(res)
	fmt.Printf("%x\n", res)
}

func TestCreatWallet(t *testing.T) {
	ws := make(Wallets)
	addr := ws.CreateWallet()
	fmt.Println(addr)
	fmt.Println(len(ws))
	fmt.Println(ws[addr].PrivateKey.D)
	ws.SaveToFile()
}
func TestWallets_LoadFromFile(t *testing.T) {
	ws := make(Wallets)
	err := ws.LoadFromFile()
	if err != nil {
		return
	}
	fmt.Println(ws["1CtRJ9GD91f6m14U3Seeg4XkXnVhXNYEAV"].PublicKey)
}

func TestWallets_addWallet(t *testing.T) {
	ws, err := NewWallets()
	if err != nil {
		return
	}
	ws.CreateWallet()

	fmt.Println(ws.GetAddresses())
}
