package heap

import (
	"testing"
)

func TestHEAP_DECREASE_KEY(t *testing.T) {
	var tests = []struct {
		a    []int
		i    int
		key  int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16},
			8,
			0,
			[]int{0, 1, 3, 2, 7, 9, 10, 14, 4, 16},
		},
		{
			[]int{16},
			0,
			15,
			[]int{15},
		},
	}

	for _, test := range tests {
		b := make([]int, len(test.a))
		copy(b, test.a)
		if HEAP_DECREASE_KEY(b, test.i, test.key); !testEq(b, test.want) {
			t.Errorf("HEAP_DECREASE_KEY(%v, %d, %d) return (%v), want (%v)", test.a, test.i, test.key, b, test.want)
		}
	}
}
