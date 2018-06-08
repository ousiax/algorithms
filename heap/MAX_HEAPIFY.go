package heap

// MAX-HEAPIFY(A, i)
//   l = LEFT(i)
//   r = RIGHT(i)
//   if l <= A.heap-size and A[l] > A[i]
//     largest = l
//   else
//     largest = r
//   if r <= A.heap-size and A[r] > A[largest]
//     largest = r
//   if largest != i
//     exchange A[i] with A[largest]
//     MAX-HEAPIFY(A, largest)
func MAX_HEAPIFY(A []int, i int) {
	l := LEFT(i)
	r := RIGHT(i)
	largest := i

	if l <= len(A)-1 && A[l] > A[i] {
		largest = l
	}

	if r <= len(A)-1 && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		MAX_HEAPIFY(A, largest)
	}
}
