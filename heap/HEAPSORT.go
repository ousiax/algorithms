package heap

// HEAPSORT(A)
//   BUILD-MAX-HEAP(A)
//   for i = A.length downto 2
//     exchange A[i] with A[1]
//     A.heap-size = A.heap-size - 1
//     MAX-HEAPIFY(A, 1)
func HEAPSORT(A []int) {
	BUILD_MAX_HEAP(A)
	for i := len(A) - 1; i >= 1; i-- {
		A[i], A[0] = A[0], A[i]
		MAX_HEAPIFY(A[0:i], 0)
	}
}
