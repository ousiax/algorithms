package heap

import (
	"fmt"
)

func ExampleBUILD_MAX_HEAP() {
	A := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Printf("INIT: %v\n", A)

	BUILD_MAX_HEAP(A)

	fmt.Printf("FIN: %v\n", A)
	// Output:
	// INIT: [4 1 3 2 16 9 10 14 8 7]
	// FIN: [16 14 10 8 7 9 3 2 4 1]
}
