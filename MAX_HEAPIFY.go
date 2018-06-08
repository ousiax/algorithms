package main

import (
	"fmt"

	. "./heap"
)

func main() {
	A := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

	fmt.Printf("INIT: %v\n", A)

	MAX_HEAPIFY(A, 1)

	fmt.Printf("FIN: %v\n", A)
}
