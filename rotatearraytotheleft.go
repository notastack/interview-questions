package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//user sets up variables
	fmt.Println("What's the size of the array?")
	var n, max int
	fmt.Scan(&n)
	A := make([]int, n, 100)
	fmt.Println("what should the max number be?")
	fmt.Scan(&max)
	//randomise the output
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		A[i] = rand.Intn(max + 1)
	}
	//prints out the original so you can check it was actually scrambled before
	fmt.Println("Original array:")
	fmt.Println(A, n)
	var l int
	fmt.Println("How many rotations should occure?")
	fmt.Scan(&l)
	//prints the results
	fmt.Println("Rotated array:")
	fmt.Println(rotateL(A, l))

}

func rotateL(a []int, i int) []int {
	//rotate the array to the left * i
	i = i % len(a)
	if i < 0 {
		i += len(a)
	}
	for c := 0; c < gcd(i, len(a)); c++ {
		t := a[c]
		j := c
		for {
			k := j + i
			if k >= len(a) {
				k -= len(a)
			}
			if k == c {
				break
			}
			a[j] = a[k]
			j = k
		}
		a[j] = t
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
