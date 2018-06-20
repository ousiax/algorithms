package stack

import (
	"testing"
)

func TestEMPTY(t *testing.T) {
	s := Stack{}
	if e := s.EMPTY(); !e {
		t.Errorf("s.EMPTY() return %t want %t", e, !e)
	}
}

func TestTOP(t *testing.T) {
	s := Stack{}
	if _, err := s.TOP(); err == nil {
		t.Errorf("s.TOP() return (_, %v) want (_, %v)", err, nil)
	}

	s.PUSH(11)
	if x, _ := s.TOP(); x != 11 {
		t.Errorf("s.TOP() return (%d, _) want (%d, _)", x, 11)
	}
}

func TestPOP(t *testing.T) {
	s := Stack{}
	if _, err := s.POP(); err == nil {
		t.Errorf("s.POP() return (_, %v) want (_, %v)", err, nil)
	}

	x := 11
	s.PUSH(x)
	if y, _ := s.POP(); y != x {
		t.Errorf("s.POP() return (%d, _) want (%d, _)", y, x)
	}
}
