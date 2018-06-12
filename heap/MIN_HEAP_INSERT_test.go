package heap

import (
	"testing"
)

func TestMIN_HEAP_INSERT(t *testing.T) {
	var tests = []struct {
		a    []int
		key  int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16},
			0,
			[]int{0, 1, 3, 4, 2, 9, 10, 14, 8, 16, 7},
		},
		{
			[]int{16},
			18,
			[]int{16, 18},
		},
	}

	for _, test := range tests {
		b := make([]int, len(test.a))
		copy(b, test.a)
		if b := MIN_HEAP_INSERT(b, test.key); !testEq(b, test.want) {
			t.Errorf("MIN_HEAP_INSERT(%v, %d) return %v, want %v", test.a, test.key, b, test.want)
		}
	}
}
