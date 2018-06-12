package heap

// HEAP-EXTRACT-MINIMUM(A)
//   if A.heap-size < 1
//     error "heap underflow"
//   max = A[1]
//   A[1] = A[A.heap-size]
//   A.heap-size = A.heap-size - 1
//   MIN-HEAPIFY(A, 1)
func HEAP_EXTRACT_MINIMUM(A []int) ([]int, int) {
	A[0], A[len(A)-1] = A[len(A)-1], A[0]
	MIN_HEAPIFY(A[0:len(A)-1], 0)
	return A[0 : len(A)-1], A[len(A)-1]
}
