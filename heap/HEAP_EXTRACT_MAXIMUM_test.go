package heap

import (
	"testing"
)

func TestHEAP_EXTRACT_MAXIMUM(t *testing.T) {
	var tests = []struct {
		a       []int
		wantB   []int
		wantMax int
	}{
		{
			[]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
			[]int{14, 8, 10, 4, 7, 9, 3, 2, 1},
			16,
		},
		{
			[]int{16},
			[]int{},
			16,
		},
	}

	for _, test := range tests {
		if b, m := HEAP_EXTRACT_MAXIMUM(test.a); !testEq(b, test.wantB) || m != test.wantMax {
			t.Errorf("HEAP_EXTRACT_MAXIMUM(%v) return (%v,%d), want (%v,%d)", test.a, b, m, test.wantB, test.wantMax)
		}
	}
}
