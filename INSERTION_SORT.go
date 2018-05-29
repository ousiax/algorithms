package main

import (
	"fmt"
)

func main() {
	A := []int{5, 2, 4, 6, 1, 3}

	fmt.Printf("INIT: %v\n", A)

	INSERTION_SORT(A)

	fmt.Printf("FIN: %v\n", A)
}

func INSERTION_SORT(A []int) {
	j := 1
	for j < len(A) {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
		j = j + 1
		fmt.Printf("ROUND %d: %v\n", j-1, A)
	}
}
