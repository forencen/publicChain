package wallet

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	a := []byte("1234")
	fmt.Println(Base58Encode(a))
}
