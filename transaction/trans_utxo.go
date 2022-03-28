package transaction

import "fmt"

type Utxo struct {
	TxHash    []byte
	Vout      int
	Value     int64
	ScriptSig string
	IsUsed    bool
}

func NewUtxo(hash []byte, vout int, value int64, scriptSig string, isUsed bool) *Utxo {
	return &Utxo{hash, vout, value, scriptSig, isUsed}
}

func (utxo *Utxo) Used() {
	utxo.IsUsed = true
}

func (utxo *Utxo) Index() string {
	return fmt.Sprintf("%x_%d", utxo.TxHash, utxo.Vout)
}
