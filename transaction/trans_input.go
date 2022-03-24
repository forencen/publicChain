package transaction

import "fmt"

type TxInput struct {
	TxHash    []byte
	Vout      int
	ScriptSig string
}

func (vin *TxInput) String() string {
	return fmt.Sprintf("%x: %d, %s", vin.TxHash, vin.Vout, vin.ScriptSig)
}
