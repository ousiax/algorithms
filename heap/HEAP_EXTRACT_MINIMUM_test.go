package heap

import (
	"testing"
)

func TestHEAP_EXTRACT_MINIMUM(t *testing.T) {
	var tests = []struct {
		a       []int
		wantB   []int
		wantMin int
	}{
		{
			[]int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16},
			[]int{2, 4, 3, 8, 7, 9, 10, 14, 16},
			1,
		},
		{
			[]int{16},
			[]int{},
			16,
		},
	}

	for _, test := range tests {
		if b, m := HEAP_EXTRACT_MINIMUM(test.a); !testEq(b, test.wantB) || m != test.wantMin {
			t.Errorf("HEAP_EXTRACT_MAXIMUM(%v) return (%v,%d), want (%v,%d)", test.a, b, m, test.wantB, test.wantMin)
		}
	}
}
