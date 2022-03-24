package pow

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"publicChain/utils"
)

type ProofOfWork struct {
	block  BlockInterface
	target *big.Int
}

const targetBit = 16

type BlockInterface interface {
	DataForPow() []byte
	SetHash([]byte)
	SetNonce(int64)
	GetHash() []byte
}

func NewProofOfWork(block BlockInterface) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) prepareDate(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.block.DataForPow(),
		utils.IntToBytes(nonce),
	}, []byte{})
	return data
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
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
	pow.block.SetHash(hash[:])
	pow.block.SetNonce(nonce)
	return hash[:], nonce
}

func (pow *ProofOfWork) IsValida() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.block.GetHash())
	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}
