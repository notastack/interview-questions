package main

import (
	"fmt"
	"time"
)

func main() {
	var cfizz, cbuzz, lower, higher int
	fmt.Println("What should the value of fizz be?")
	fmt.Scanln(&cfizz)
	fmt.Println("What should the value of buzz be?")
	fmt.Scanln(&cbuzz)
	fmt.Println("Where should the counting start?")
	fmt.Scanln(&lower)
	fmt.Println("Where should the counting stop?")
	fmt.Scanln(&higher)
	fibu(cfizz, cbuzz, lower, higher)

}

func fibu(cfizz, cbuzz, lower, higher int) {
	for i := lower; i < higher+1; i++ {
		switch {
		case i%(cfizz*cbuzz) == 0:
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
