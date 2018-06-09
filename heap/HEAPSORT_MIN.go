package heap

// HEAPSORT-MIN(A)
//   BUILD-MIN-HEAP(A)
//   for i=A.heap-size downto 2
//     exchange A[1] with A[i]
//     A.heap-size = A.heap-size - 1
//     MIN-HEAPIFY(A, 1)
func HEAPSORT_MIN(A []int) {
	BUILD_MIN_HEAP(A)
	for i := len(A) - 1; i > 0; i-- {
		A[i], A[0] = A[0], A[i]
		MIN_HEAPIFY(A[0:i], 0)
	}
}
