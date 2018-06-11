package heap

import (
	"testing"
)

func TestHEAP_INCREASE_KEY(t *testing.T) {
	var tests = []struct {
		a    []int
		i    int
		key  int
		want []int
	}{
		{
			[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
			8,
			15,
			[]int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1},
		},
		{
			[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1, 0},
			10,
			15,
			[]int{16, 15, 10, 8, 14, 9, 3, 2, 4, 1, 7},
		},
		{
			[]int{16},
			0,
			17,
			[]int{17},
		},
	}

	for _, test := range tests {
		a := make([]int, len(test.a))
		copy(a, test.a)
		if HEAP_INCREASE_KEY(a, test.i, test.key); !testEq(a, test.want) {
			t.Errorf("HEAP_INCREASE_KEY(%v, %d, %d) return %v, want %v", test.a, test.i, test.key, a, test.want)
		}
	}
}
