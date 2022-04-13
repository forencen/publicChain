package transaction

import "fmt"

type Utxo struct {
	TxHash    []byte
	Vout      int
	Value     int64
	ScriptSig string
}

func NewUtxo(hash []byte, vout int, value int64, scriptSig string) *Utxo {
	return &Utxo{hash, vout, value, scriptSig}
}

func (utxo *Utxo) Index() string {
	return fmt.Sprintf("%x_%d", utxo.TxHash, utxo.Vout)
}
