package main

import (
	"fmt"
	"time"
)

func main() {

	cfizz := 3
	cbuzz := 5
	for i := 1; i < 101; i++ {
		switch {
		case i%(cfizz * cbuzz) == 0:
			fmt.Println("fizzbuzz")
		case i%cfizz == 0:
			fmt.Println("fizz")
		case i%cbuzz == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
		time.Sleep(1 * time.Second)
	}

}
