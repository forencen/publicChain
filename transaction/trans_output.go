package transaction

import (
	"bytes"
	"fmt"
	"publicChain/wallet"
)

type TxOutput struct {
	Value      int64
	PubKeyHash []byte
}

func (vout *TxOutput) Lock(address string) {
	vout.PubKeyHash = wallet.GetAddressPubKeyHash(address)
}

func (vout *TxOutput) UnLockWithAddress(pubKeyHash []byte) bool {
	return bytes.Compare(pubKeyHash, vout.PubKeyHash) == 0
}

func (vout *TxOutput) String() string {
	return fmt.Sprintf("%d, %s", vout.Value, wallet.TransformPublicKeyHash(vout.PubKeyHash))
}
