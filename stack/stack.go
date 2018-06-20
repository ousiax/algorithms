package stack

import (
	"errors"
)

type Stack struct {
	s   []int
	top int
}

// STACK-EMPTY(S)
//   if S.top == 0
//     return TRUE
//   else
//     return FALSE
func (s *Stack) EMPTY() bool {
	return s.top == 0
}

// PUSH(S,x)
//   S.top = S.top + 1
//   S[S.top] = x
func (s *Stack) PUSH(x int) {
	s.s = append(s.s, x)
	s.top += 1
}

// POP(S)
//   if STACK-EMPTY(S)
//     error "underflow"
//   else
//     S.top = S.top - 1
//     return S[S.top + 1]
func (s *Stack) POP() (x int, err error) {
	if s.EMPTY() {
		return 0, errors.New("underflow")
	}

	s.top -= 1
	return s.s[s.top], nil
}

// TOP(S)
//   if STACK-EMPTY(S)
//     error "underflow"
//   else
//     return S[S.top]
func (s *Stack) TOP() (x int, err error) {
	if s.EMPTY() {
		return 0, errors.New("underflow")
	}

	return s.s[s.top-1], nil
}
