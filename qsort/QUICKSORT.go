package qsort

// QUICKSORT(A, p, r)
//   if p < r
//     q = PARTITION(A, p, r)
//     QUICKSORT(A, p, q-1)
//     QUICKSORT(A, q+1, r)
func QUICKSORT(A []int, p, r int) {
	if p >= r {
		return
	}
	q := PARTITION(A, p, r)
	QUICKSORT(A, p, q-1)
	QUICKSORT(A, q+1, r)
}

// PARTITION(A, p, r)
//   x = A[r] 			// 主元(pivot element)
//   i = p - 1
//   for j = p to r-1
//     if A[j] <= x
//       i = i + 1
//       exchange A[i], A[j]
//   exchange A[i+1], A[r]
//   return i + 1
func PARTITION(A []int, p, r int) (q int) {
	x := A[r]
	i := p - 1
	// 1. 若 p<= k <= i, 则 A[k] <= x
	// 2. 若 i+1 <= k <= r-1, 则 A[k] >= x
	// 3. 若 k = r, 则 A[k] = x
	for j := p; j < r; j++ {
		if A[j] <= x {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
