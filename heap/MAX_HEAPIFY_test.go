package heap

import (
	"fmt"
)

func ExampleMAX_HEAPIFY() {
	A := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}

	fmt.Printf("INIT: %v\n", A)

	MAX_HEAPIFY(A, 1)

	fmt.Printf("FIN: %v\n", A)
	// Output:
	// INIT: [16 4 10 14 7 9 3 2 8 1]
	// FIN: [16 14 10 8 7 9 3 2 4 1]
}
