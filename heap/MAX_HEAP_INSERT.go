package heap

import (
	"math"
)

// MAX-HEAP-INSERT(A, key)
//   A.heap-size = A.heap-size + 1
//   A[A.heap-size] = int.MinValue
//   HEAP-INCREASE-KEY(A, A.heap-size, key)
func MAX_HEAP_INSERT(A []int, key int) []int {
	B := append(A, math.MinInt64)
	HEAP_INCREASE_KEY(B, len(B)-1, key)
	return B
}
