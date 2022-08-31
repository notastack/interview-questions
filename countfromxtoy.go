package main

import (
	"fmt"
)

func main() {
	var cfizz, cbuzz int
	fmt.Println("where should we start?")
	fmt.Scanln(&cfizz)
	fmt.Println("What should we end?")
	fmt.Scanln(&cbuzz)

	fibu(cfizz, cbuzz)

}

func fibu(x, y int) {
	for i := x; i < y+1; i++ {
		fmt.Println(i)

	}

}
