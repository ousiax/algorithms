package stack

import (
	"fmt"
)

func ExampleSTACK() {
	s := Stack{}
	s.PUSH(4)
	s.PUSH(1)
	s.PUSH(3)
	s.POP()
	s.PUSH(8)
	s.POP()

	for {
		x, err := s.POP()
		if err != nil {
			break
		}
		fmt.Printf("%d ", x)
	}

	fmt.Printf("\n")

	// Output:
	// 1 4
}
