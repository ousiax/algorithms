package stable_sort

import (
	"fmt"
)

func ExampleCOUNTING_SORT() {
	A := []int{2, 5, 3, 0, 2, 3, 0, 3}
	k := K(A)
	B := make([]int, len(A))

	fmt.Printf("INIT:%v\n", A)

	COUNTING_SORT(A, B, k)

	fmt.Printf("FIN:%v\n", B)
	// Output:
	// INIT:[2 5 3 0 2 3 0 3]
	// FIN:[0 0 2 2 3 3 3 5]
}
