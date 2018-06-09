package heap

// BUILD-MIN-HEAP(A)
//   A.heap-size = A.length
//   for i = [len(A) / 2] downto 1  // round down
//     MIN-HEAPIFY(A, i)
func BUILD_MIN_HEAP(A []int) {
	for i := len(A) / 2; i >= 0; i-- {
		MIN_HEAPIFY(A, i)
	}
}
