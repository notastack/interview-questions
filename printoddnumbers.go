package main

import (
	"fmt"
)

func main() {
	var cfizz, cbuzz int
	fmt.Println("Where should we start?")
	fmt.Scanln(&cfizz)
	fmt.Println("Where should we end?")
	fmt.Scanln(&cbuzz)

	fibu(cfizz, cbuzz)

}

func fibu(x, y int) {
	for i := x; i < y+1; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}

	}

}
