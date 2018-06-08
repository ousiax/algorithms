package heap

// BUILD-MAX-HEAP(A)
//   A.heap-size = A.length
//   for i = [A.length/2] downto 1  // round down
//     MAX-HEAPIFY(A, i)
func BUILD_MAX_HEAP(A []int) {
	for i := len(A) / 2; i >= 0; i-- {
		MAX_HEAPIFY(A, i)
	}
}
