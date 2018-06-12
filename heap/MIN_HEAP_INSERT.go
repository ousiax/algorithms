package heap

import (
	"math"
)

// MIN-HEAP-INSERT(A, key)
//   A.heap-size = A.heap-size + 1
//   A[A.heap-size] = int.MaxValue
//   HEAP-DECREASE-KEY(A, A.heap-size, key)
func MIN_HEAP_INSERT(A []int, key int) []int {
	B := append(A, math.MaxInt32)
	HEAP_DECREASE_KEY(B, len(B)-1, key)
	return B
}
