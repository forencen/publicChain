package wallet

import (
	"fmt"
	"math/big"
)

/*
Base58 是一种二进制转化为可视字符串的算法，主要用来转换大整数。区别是，转换出来的字符串，去除了几个看起来会产生歧义的字符
*/

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte
	// input 转化为大数字
	x := big.NewInt(0).SetBytes(input)
	// 定义进制
	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}
	fmt.Println(result)
	ReverseBytes(result)
	fmt.Println(result)
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result
}

func ReverseBytes(input []byte) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
