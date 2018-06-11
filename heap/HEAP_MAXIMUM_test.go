package heap

import (
	"testing"
)

func TestHEAP_MAXIMUM(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{
			[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
			16,
		},
		{
			[]int{16},
			16,
		},
	}

	for _, test := range tests {
		if w := HEAP_MAXIMUM(test.a); w != test.want {
			t.Errorf("HEAP_MAXIMUM(%v) return %d, want %d", test.a, w, test.want)
		}
	}
}
