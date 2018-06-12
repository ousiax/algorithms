package heap

// HEAP-DECREASE-KEY(A, i, key)
//   if key > A[i]
//     error "new key is greater than current key"
//   A[i] = key
//   for i > 1 && A[PARENT(i)] > key
//     A[i], A[PARENT(i)] = A[PARENT(i)], A[i]
//     i = PARENT(i)
func HEAP_DECREASE_KEY(A []int, i, key int) {
	if key > A[i] {
		return
	}

	A[i] = key
	for i > 0 && A[PARENT(i)] > key {
		A[i], A[PARENT(i)] = A[PARENT(i)], A[i]
		i = PARENT(i)
	}
}
