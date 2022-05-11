package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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

func (t *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTxs map[string]Transaction) {
	if t.IsCoinbase() {
		return
	}
	txTrimmed := t.TrimmedCopy()
	for vinId, vin := range txTrimmed.Vins {
		prevTx := prevTxs[hex.EncodeToString(vin.TxHash)]
		txTrimmed.Vins[vinId].Signature = nil
		txTrimmed.Vins[vinId].PubKey = prevTx.Vouts[vin.Vout].PubKeyHash
		txTrimmed.SetTxHash()
		txTrimmed.Vins[vinId].PubKey = nil

		r, s, _ := ecdsa.Sign(rand.Reader, &privateKey, txTrimmed.Hash)
		signature := append(r.Bytes(), s.Bytes()...)

		t.Vins[vinId].Signature = signature
	}
}

func (t *Transaction) TrimmedCopy() Transaction {
	var (
		inputs  []*TxInput
		outputs []*TxOutput
	)
	for _, input := range t.Vins {
		inputs = append(inputs, &TxInput{input.TxHash, input.Vout, nil, nil})
	}
	for _, output := range t.Vouts {
		outputs = append(outputs, &TxOutput{output.Value, output.PubKeyHash})
	}
	return Transaction{t.Hash, inputs, outputs}
}

func (t *Transaction) IsCoinbase() bool {
	return len(t.Vins) == 1 && len(t.Vins[0].TxHash) == 0 && t.Vins[0].Vout == -1
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
