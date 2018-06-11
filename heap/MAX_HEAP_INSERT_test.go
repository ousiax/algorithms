package heap

import (
	"testing"
)

func TestMAX_HEAP_INSERT(t *testing.T) {
	var tests = []struct {
		a    []int
		key  int
		want []int
	}{
		{
			[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
			15,
			[]int{16, 15, 10, 8, 14, 9, 3, 2, 4, 1, 7},
		},
		{
			[]int{16},
			17,
			[]int{17, 16},
		},
	}

	for _, test := range tests {
		b := make([]int, len(test.a))
		copy(b, test.a)
		if b := MAX_HEAP_INSERT(b, test.key); !testEq(b, test.want) {
			t.Errorf("MAX_HEAP_INSERT(%v, %d) return %v, want %v", test.a, test.key, b, test.want)
		}
	}
}
