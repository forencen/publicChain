package transaction

import (
	"bytes"
	"fmt"
	"publicChain/wallet"
)

type TxInput struct {
	TxHash    []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

// UnLockWithAddress  wallet 中的公钥
func (vin *TxInput) UnLockWithAddress(pubKeyHash []byte) bool {
	lockingHash := wallet.HashPubKey(vin.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (vin *TxInput) String() string {
	return fmt.Sprintf("%x: %d, %s", vin.TxHash, vin.Vout, wallet.TransformPublicKey(vin.PubKey))
}
