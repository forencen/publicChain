package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"publicChain/pow"
	"publicChain/transaction"
	"publicChain/utils"
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
	pow.NewProofOfWork(block).Run()
	return block
}

func CreateGenesisBlock(txs []*transaction.Transaction) *Block {
	preHash := [32]byte{}
	genesisBlock := &Block{
		Height: 1, Timestamp: time.Now().Unix(), PrevBlockHash: preHash[:],
		Txs: txs, Hash: []byte{}, Nonce: 0,
	}
	pow.NewProofOfWork(genesisBlock).Run()
	return genesisBlock
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	if err := encoder.Encode(b); err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}

func (b *Block) SetHash(h []byte) {
	b.Hash = h
}

func (b *Block) GetHash() []byte {
	return b.Hash
}

func (b *Block) SetNonce(nonce int64) {
	b.Nonce = nonce
}

func (b *Block) DataForPow() []byte {
	return bytes.Join([][]byte{
		b.PrevBlockHash,
		b.HashTransactions(),
		utils.IntToBytes(b.Timestamp),
		utils.IntToBytes(b.Height),
	}, []byte{})
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

func (b *Block) PrintBlock() {
	fmt.Println("---------------------------------------")
	fmt.Printf("block height:%d\n", b.Height)
	fmt.Printf("block hash:%x\n", b.Hash)
	fmt.Printf("block pre hash:%x\n", b.PrevBlockHash)
	fmt.Println("block txs:")
	for _, tx := range b.Txs {
		fmt.Println(tx)
	}
	fmt.Printf("%s\n", time.Unix(b.Timestamp, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("---------------------------------------")
}

func Deserialize(blockBytes []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&b); err != nil {
		log.Panic(err)
	}
	return &b
}
