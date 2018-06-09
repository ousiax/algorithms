package heap

// MIN_HEAPIFY(A, i)
//   l = LEFT(i)
//   r = RIGHT(i)
//   least = i
//   if r < A.heap-size && A[l] < A[least]
//     least = l
//   if r < A.heap-size && A[r] < A[least]
//     least = r
//   if least != i
//     exchange A[least] with A[i]
//     MIN_HEAPIFY(A, least)
func MIN_HEAPIFY(A []int, i int) {
	l := LEFT(i)
	r := RIGHT(i)
	least := i

	if l < len(A) && A[l] < A[least] {
		least = l
	}

	if r < len(A) && A[r] < A[least] {
		least = r
	}

	if least != i {
		A[least], A[i] = A[i], A[least]
		MIN_HEAPIFY(A, least)
	}
}
