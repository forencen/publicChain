package wallet

import (
	"fmt"
	"math/big"
	"testing"
)

func TestEncode(t *testing.T) {
	a := []byte("0001234")
	fmt.Println(a)
	fmt.Println(big.NewInt(0).SetBytes(a))
	fmt.Println(Base58Encode(a))
	fmt.Println(Base58Decode(Base58Encode(a)))
}
