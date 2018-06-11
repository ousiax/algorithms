package heap

// HEAP-INCREASE-KEY(A, i, key)
//   if key < A[i]
//     error "new key is smaller than current key"
//   A[i] = key
//   while i > 1 and A[PARENT(i)] < A[i]
//     exchange A[i] with A[PARENT(i)]
//     i = PARENT(i)
func HEAP_INCREASE_KEY(A []int, i, key int) {
	A[i] = key
	for i > 0 && A[PARENT(i)] < A[i] {
		A[i], A[PARENT(i)] = A[PARENT(i)], A[i]
		i = PARENT(i)
	}
}

// func HEAP_INCREASE_KEY(A []int, i, key int) {
// 	if A[i] > key {
// 		return
// 	}
//
// 	A[i] = key
// 	p := PARENT(i)
// 	if p >= 0 && A[p] < A[i] {
// 		A[i] = A[p]
// 		if p > 0 {
// 			HEAP_INCREASE_KEY(A, p, key)
// 		}
// 	}
// }
