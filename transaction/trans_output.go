package transaction

import "fmt"

type TxOutput struct {
	Value     int64
	ScriptSig string
}

func (vout *TxOutput) UnLockWithAddress(address string) bool {
	return vout.ScriptSig == address
}

func (vout *TxOutput) String() string {
	return fmt.Sprintf("%d, %s", vout.Value, vout.ScriptSig)
}
