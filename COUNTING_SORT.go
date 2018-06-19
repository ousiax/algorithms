package main

import (
	"fmt"

	. "./stable_sort"
)

func main() {
	A := []int{6, 0, 2, 0, 1, 3, 4, 6, 1, 3, 2}
	k := K(A)
	B := make([]int, len(A))

	fmt.Printf("INIT:%v\n", A)

	COUNTING_SORT(A, B, k)

	fmt.Printf("FIN:%v\n", B)
	// Output:
	// INIT:[6 0 2 0 1 3 4 6 1 3 2]
	// FIN:[0 0 1 1 2 2 3 3 4 6 6]
}
