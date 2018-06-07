package main

import (
	"fmt"
)

func main() {
	A := []int{5, 2, 4, 7, 1, 3, 2, 6}

	fmt.Printf("INIT: %v\n", A)

	MERGE_SORT(A, 0, len(A)-1)

	fmt.Printf("FIN: %v\n", A)
}

// MERGER_SORT(A,p,r)
//   if p < r
//     q = [(p+r)/2] // round down
//     MERGE_SORT(A,p,q)
//     MERGE_SORT(A,q+1,r)
//     MERGE(A,p,r)
func MERGE_SORT(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MERGE_SORT(A, p, q)
		MERGE_SORT(A, q+1, r)
		MERGE(A, p, q, r)
	}
}

// MERGE(A,p,q,r)
//
// n1 = q - p + 1
// n2 = r - q
// let L[1...n1+1] and R[1..n2+1] be new arrays
// for i = 1 to n1
//     L[i] = A[q+i]
// for j = 1 to n2
//     R[j] = A[q+j+1]
// L[n1+1] = L[n1] + R[n2]
// R[n2+1] = L[n1] + R[n2]
// i = 1
// j = 1
// for k = p to r
//     if L[i] <= R[j]
//         A[k] = L[i]
//         i = i + 1
//     else
//         A[k] = R[j]
//         j = j + 1
func MERGE(A []int, p, q, r int) {
	N1 := q - p + 1
	N2 := r - q

	L := make([]int, N1+1)
	R := make([]int, N2+1)

	for i := 0; i < N1; i++ {
		L[i] = A[p+i]
	}

	for i := 0; i < N1; i++ {
		R[i] = A[q+i+1]
	}

	L[N1] = L[N1-1] + R[N2-1]
	R[N2] = L[N1]

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i += 1
		} else {
			A[k] = R[j]
			j += 1
		}
	}
}
