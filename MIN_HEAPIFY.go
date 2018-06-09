package main

import (
	"fmt"

	. "./heap"
)

func main() {
	A := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

	fmt.Printf("INIT: %v\n", A)

	MIN_HEAPIFY(A, 0)

	fmt.Printf("FIN: %v\n", A)

	// Output:
	// INIT: [16 4 10 14 7 9 3 2 8 1]
	// FIN: [4 7 10 14 1 9 3 2 8 16]
}
