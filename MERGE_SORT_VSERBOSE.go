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

func MERGE_SORT(A []int, p, r int) {
	if p < r {
		q := (r + p) / 2
		MERGE_SORT(A, p, q)
		MERGE_SORT(A, q+1, r)
		MERGE(A, p, q, r)
	}
}

func MERGE(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}

	L[n1] = L[n1-1] + R[n2-1]
	R[n2] = L[n1]

	fmt.Printf("p=%d, q=%d, r=%d, A: %v, L: %v, R: %v\n", p, q, r, A, L, R)

	i := 0
	j := 0
	for k := p; k <= r; k++ {
		fmt.Printf("\ti=%d, j=%d, k=%d, A?: %v\n", i, j, k, A)
		if L[i] <= R[j] {
			A[k] = L[i]
			i = i + 1
		} else {
			A[k] = R[j]
			j = j + 1
		}
		fmt.Printf("\t\ti=%d, j=%d, k=%d, A?: %v\n", i, j, k, A)
	}
}
