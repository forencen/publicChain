package transaction

import "fmt"

type TxOutput struct {
	value     int64
	ScriptSig string
}

func (vout *TxOutput) String() string {
	return fmt.Sprintf("%d, %s", vout.value, vout.ScriptSig)
}
