package transaction

type Utxo struct {
	TxHash []byte
	Vout   int
	Value  int64
	PubKey []byte
}

func NewUtxo(hash []byte, vout int, value int64, pubKey []byte) *Utxo {
	return &Utxo{hash, vout, value, pubKey}
}
