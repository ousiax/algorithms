package heap

import (
	"testing"
)

func TestHEAP_MINIMUM(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{
			[]int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16},
			1,
		},
		{
			[]int{1},
			1,
		},
	}

	for _, test := range tests {
		if w := HEAP_MINIMUM(test.a); w != test.want {
			t.Errorf("HEAP_MINIMUM(%v) return %d, want %d", test.a, w, test.want)
		}
	}
}
