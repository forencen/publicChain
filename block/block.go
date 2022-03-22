package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"publicChain/transaction"
	"time"
)

type Block struct {
	Height        int64
	Nonce         int64
	Timestamp     int64
	PrevBlockHash []byte
	Txs           []*transaction.Transaction
	Hash          []byte
}

func NewBlock(height int64, txs []*transaction.Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Height: height, Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash,
		Txs: txs, Hash: []byte{}, Nonce: 0,
	}
	return block
}

func CreateGenesisBlock(txs []*transaction.Transaction) *Block {
	preHash := [32]byte{}
	return NewBlock(1, txs, preHash[:])
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	if err := encoder.Encode(b); err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}

// HashTransactions 提供挖矿时使用
// 把交易hash拼接，并且生产hash
func (b *Block) HashTransactions() []byte {
	var (
		txHashes [][]byte
		// 1byte = 8bit
		txHash [32]byte
	)
	for _, tx := range b.Txs {
		txHashes = append(txHashes, tx.Hash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func Deserialize(blockBytes []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&b); err != nil {
		log.Panic(err)
	}
	return &b
}
