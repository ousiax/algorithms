package main

import (
	"fmt"
)

func main() {
	A := []int{5, 2, 4, 6, 1, 3}

	fmt.Printf("INIT: %v\n", A)

	BUBBLE_SORT(A)

	fmt.Printf("FIN: %v\n", A)
}

func BUBBLE_SORT(A []int) {
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
		fmt.Printf("ROUND %d: %v\n", i, A)
	}
}
