package main

import (
	"fmt"
)

func main() {
	var cfizz, cbuzz, lower int
	//user sets variables
	fmt.Println("What number should we count down from?")
	fmt.Scanln(&cfizz)
	fmt.Println("Untill we reach what number?")
	fmt.Scanln(&cbuzz)
	fmt.Println("What should the decrements be?")
	fmt.Scanln(&lower)

	fibu(cfizz, cbuzz, lower)

}

func fibu(x, y, z int) {
	var i int
	i = x
	//count in z decrements down untill the number is under the limit
	for i > y {
		fmt.Println(i)
		i = i - z
	}

}
