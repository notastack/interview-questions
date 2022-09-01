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
	//prints the results
	fmt.Println("Sorted array:")
	fmt.Println(shellsort(A, n))

}

func shellsort(A []int, n int) []int {
	//go over the array over and over again untill it's sorted
	for interval := int(n / 2); interval > 0; interval /= 2 {
		for j := interval; j < n; j++ {
			for k := j; k >= interval && A[k-interval] > A[k]; k -= interval {
				A[k], A[k-interval] = A[k-interval], A[k]
			}
		}
	}
	//returns the sorted array
	return A

}
