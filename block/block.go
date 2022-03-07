package block

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Height        int64
	Nonce         int64
	Timestamp     int64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
}

func NewBlock(height int64, data string, prevBlockHash []byte) *Block {
	block := &Block{
		Height: height, Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash,
		Data: []byte(data), Hash: []byte{}, Nonce: 0,
	}
	return block
}

func CreateGenesisBlock(data string) *Block {
	preHash := [32]byte{}
	return NewBlock(1, data, preHash[:])
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	if err := encoder.Encode(b); err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}

func Deserialize(blockBytes []byte) *Block {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&b); err != nil {
		log.Panic(err)
	}
	return &b
}
