package main

import (
	"fmt"

	. "./qsort"
)

func main() {
	A := []int{2, 8, 7, 1, 3, 5, 6, 4}

	fmt.Printf("INIT: %v\n", A)

	QUICKSORT(A, 0, len(A)-1)

	fmt.Printf("FIN: %v\n", A)
}
