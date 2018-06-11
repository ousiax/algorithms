package heap

// HEAP-EXTRACT-MAXIMUM(A)
//   if A.heap-size < 1
//     error "heap underflow"
//   max = A[1]
//   A[1] = A[A.heap-size]
//   A.heap-size = A.heap-size - 1
//   MAX-HEAPIFY(A, 1)
func HEAP_EXTRACT_MAXIMUM(A []int) ([]int, int) {
	A[0], A[len(A)-1] = A[len(A)-1], A[0]
	MAX_HEAPIFY(A[0:len(A)-1], 0)
	return A[0 : len(A)-1], A[len(A)-1]
}
