package main

import (
	"fmt"
)

func main() {
	var cfizz, result int
	//user sets variables
	fmt.Println("Which number do you want the catalan number equilavent to?")
	fmt.Scanln(&cfizz)
	result = catalan(cfizz)
	fmt.Println(result)
}

func catalan(fizz int) int {
	var result int
	result = (int(IterativeFactorial(2 * fizz))) / (int(IterativeFactorial(fizz+1)) * int(IterativeFactorial(fizz)))
	return result
}

func IterativeFactorial(number int) uint64 {
	var result uint64 = 1
	if number < 0 {

	} else {
		for i := 1; i <= number; i++ {
			result *= uint64(i)

		}
	}
	return result
}
