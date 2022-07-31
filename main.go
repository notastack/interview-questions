package main

import (
	"fmt"
	"time"
)

func main() {
	var cfizz, cbuzz, lower, higher int
	//set up is made by the user himself
	fmt.Println("What should the value of fizz be?")
	fmt.Scanln(&cfizz)
	fmt.Println("What should the value of buzz be?")
	fmt.Scanln(&cbuzz)
	fmt.Println("Where should the counting start?")
	fmt.Scanln(&lower)
	fmt.Println("Where should the counting stop?")
	fmt.Scanln(&higher)
	//launch the fizzbuzz function
	fibu(cfizz, cbuzz, lower, higher)

}

func fibu(cfizz, cbuzz, lower, higher int) {
	//loop from lower end to higher end
	for i := lower; i < higher+1; i++ {
		switch {
		//if divisible by f and b then always divisible by f*b
		case i%(cfizz*cbuzz) == 0:
			fmt.Println("fizzbuzz")
		case i%cfizz == 0:
			fmt.Println("fizz")
		case i%cbuzz == 0:
			fmt.Println("buzz")
		//error is default state so not divisible by f or b
		default:
			fmt.Println(i)
		}
		//give you time to read the answer
		time.Sleep(1 * time.Second)
	}

}
