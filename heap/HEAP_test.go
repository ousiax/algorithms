package heap

import (
	"testing"
)

func TestPARENT(t *testing.T) {
	var tests = []struct {
		i    int
		want int
	}{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 3},
	}

	for _, test := range tests {
		if w := PARENT(test.i); w != test.want {
			t.Errorf("PARENT(%d) returned %d, want %d", test.i, w, test.want)
		}
	}
}

func TestLEFT(t *testing.T) {
	var tests = []struct {
		i    int
		want int
	}{
		{0, 1},
		{1, 3},
		{2, 5},
		{3, 7},
		{4, 9},
		{5, 11},
	}

	for _, test := range tests {
		if w := LEFT(test.i); w != test.want {
			t.Errorf("LEFT(%d) returned %d, want %d", test.i, w, test.want)
		}
	}
}

func TestRIGHT(t *testing.T) {
	var tests = []struct {
		i    int
		want int
	}{
		{0, 2},
		{1, 4},
		{2, 6},
		{3, 8},
		{4, 10},
		{5, 12},
	}

	for _, test := range tests {
		if w := RIGHT(test.i); w != test.want {
			t.Errorf("RIGHT(%d) returned %d, want %d", test.i, w, test.want)
		}
	}
}
