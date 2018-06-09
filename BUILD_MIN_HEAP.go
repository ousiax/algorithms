package main

import (
	"fmt"

	. "./heap"
)

func main() {
	A := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Printf("INIT: %v\n", A)

	BUILD_MIN_HEAP(A)

	fmt.Printf("FIN: %v\n", A)

	// Output:
	// INIT: [4 1 3 2 16 9 10 14 8 7]
	// FIN: [1 2 3 4 7 9 10 14 8 16]
}
