package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {
	Hash  []byte
	Vins  []*TxInput
	Vouts []*TxOutput
}

func (t *Transaction) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(t)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func (t *Transaction) SetTxHash() {
	hash := sha256.Sum256(t.Serialize())
	t.Hash = hash[:]
}

func NewCoinbaseTransaction(address string) *Transaction {
	txInput := &TxInput{[]byte{}, -1, "genesis data"}
	txOutput := &TxOutput{10, address}
	coinbase := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	coinbase.SetTxHash()
	return coinbase
}
