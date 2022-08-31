package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("What's the size of the array?")
	var n, max int
	fmt.Scan(&n)
	A := make([]int, n, 100)
	fmt.Println("what should the max number be?")
	fmt.Scan(&max)
	for i := 0; i < n; i++ {
		A[i] = rand.Intn(max + 1)
	}
	fmt.Println("Original array:")
	fmt.Println(A, n)
	fmt.Println("Sorted array:")
	fmt.Println(shellsort(A, n))

}

func shellsort(A []int, n int) []int {

	for interval := int(n / 2); interval > 0; interval /= 2 {
		for j := interval; j < n; j++ {
			for k := j; k >= interval && A[k-interval] > A[k]; k -= interval {
				A[k], A[k-interval] = A[k-interval], A[k]
			}
		}
	}
	return A

}
