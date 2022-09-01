package main

import (
	"fmt"
)

func main() {
	var cfizz, cbuzz, lower int
	//user sets up variables
	fmt.Println("Where should we start?")
	fmt.Scanln(&cfizz)
	fmt.Println("Where should we end?")
	fmt.Scanln(&cbuzz)
	fmt.Println("What number should I be divisible by?")
	fmt.Scanln(&lower)

	fibu(cfizz, cbuzz, lower)

}

func fibu(x, y, z int) {
	for i := x; i < y+1; i++ {
		//check if i is divisible by z
		if i%z == 0 { 
			fmt.Println(i)
		}
		
	}

}
