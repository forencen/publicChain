package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Wallets map[string]*Wallet

func (w Wallets) SaveToFile() {
	var content bytes.Buffer
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(w)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

func (w Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}
	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}
	var tempWallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	if err = decoder.Decode(&tempWallets); err != nil {
		return err
	}
	for key, value := range tempWallets {
		w[key] = value
	}
	return nil
}

func (w Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())
	w[address] = wallet
	return address
}

func (w Wallets) GetAddresses() []string {
	var addresses []string
	for address := range w {
		addresses = append(addresses, address)
	}
	return addresses
}

func (w Wallets) GetWallet(address string) *Wallet {
	return w[address]
}

func NewWallets() (Wallets, error) {
	wallets := make(Wallets)
	err := wallets.LoadFromFile()
	return wallets, err
}
