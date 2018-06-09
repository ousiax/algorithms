package heap

import (
	"fmt"
)

func ExampleHEAPSORT_MIN() {
	A := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Printf("INIT: %v\n", A)

	HEAPSORT_MIN(A)

	fmt.Printf("FIN: %v\n", A)

	// Output:
	// INIT: [4 1 3 2 16 9 10 14 8 7]
	// FIN: [16 14 10 9 8 7 4 3 2 1]
}
