package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"publicChain/wallet"
	"strings"
)

const Subsidy = 10

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

func (t *Transaction) TrimmedCopy() {

}

func (t *Transaction) String() string {
	var sBuilder strings.Builder
	sBuilder.WriteString("交易hash:")
	sBuilder.WriteString(fmt.Sprintf("%x\n", t.Hash))
	sBuilder.WriteString("vin:\n")
	for _, vin := range t.Vins {
		sBuilder.WriteString(fmt.Sprintf("%s\n", vin))
	}
	sBuilder.WriteString("vout:\n")
	for _, vout := range t.Vouts {
		sBuilder.WriteString(fmt.Sprintf("%s\n", vout))
	}
	return sBuilder.String()
}

func NewCoinbaseTransaction(address string) *Transaction {
	txInput := &TxInput{[]byte{}, -1, nil, []byte("genesis coinbase")}
	txOutput := &TxOutput{Subsidy, wallet.GetAddressPubKeyHash(address)}
	coinbase := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	coinbase.SetTxHash()
	return coinbase
}
