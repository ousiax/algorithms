package qsort

import (
	"fmt"
	"testing"
)

func ExampleQUICKSORT() {
	A := []int{49, 38, 65, 97, 76, 13, 27, 49}
	QUICKSORT(A, 0, len(A)-1)
	fmt.Printf("%v\n", A)
	// Output:
	// [13 27 38 49 49 65 76 97]
}
