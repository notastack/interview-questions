package main

import (
	"fmt"
	"math/big"
)

func main() {
	var cfizz int
	cfizz = 0
	//user sets variables
	fmt.Println("Which number do you want the catalan number equilavent to?")
	fmt.Scanln(&cfizz)
	catalan(cfizz)
}

func catalan(n int) {
	var b, c big.Int
	v := int64(n)
	fmt.Println("the catalan number for ", n, " is ", (c.Div(b.Binomial(v*2, v), c.SetInt64(v+1))), " .")
}
