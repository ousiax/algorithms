package main

import (
	"fmt"
)

func main() {
	A := []int{5, 2, 4, 7, 1, 3, 2, 6}

	fmt.Printf("INIT: %v\n", A)

	MERGER_SORT(A, 0, len(A)-1)

	fmt.Printf("FIN: %v\n", A)
}

func MERGER_SORT(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MERGER_SORT(A, p, q)
		MERGER_SORT(A, q+1, r)
		MERGER(A, p, q, r)
	}
}

func MERGER(A []int, p, q, r int) {
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
