package BLG

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

const targetBit = 16

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) prepareDate(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevBlockHash,
		pow.Block.Data,
		IntToBytes(pow.Block.Timestamp),
		IntToBytes(nonce),
		IntToBytes(pow.Block.Height),
	}, []byte{})
	return data
}

func (pow *ProofOfWork) run() ([]byte, int64) {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt big.Int

	for {
		dataBytes := pow.prepareDate(nonce)
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	return hash[:], nonce
}

func (pow *ProofOfWork) IsValida() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}
