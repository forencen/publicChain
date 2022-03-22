package transaction

type TxInput struct {
	TxHash    []byte
	Vout      int
	ScriptSig string
}
