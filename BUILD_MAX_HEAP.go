package main

import (
	"fmt"

	. "./heap"
)

func main() {
	A := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Printf("INIT: %v\n", A)

	BUILD_MAX_HEAP(A)

	fmt.Printf("FIN: %v\n", A)
}
