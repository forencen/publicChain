package main

import (
	"fmt"
	"math/big"
)

func main() {
	//c := &cli.Cli{}
	//c.Run()
	input := []byte("1234")
	fmt.Println(big.NewInt(0).SetBytes(input))
	a := big.NewInt(5)
	d := big.NewInt(6)
	b := big.NewInt(2)
	c := &big.Int{}
	a.DivMod(d, b, c)
	fmt.Println(c)
}
