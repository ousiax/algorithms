package bst

import (
	"fmt"
)

func ExampleBST() {
	t := &Tree{}
	t.INSERT(NewNode(15))
	t.INSERT(NewNode(4))
	t.INSERT(NewNode(2))
	t.INSERT(NewNode(13))
	t.INSERT(NewNode(9))
	t.INSERT(NewNode(18))
	t.INSERT(NewNode(17))
	t.INSERT(NewNode(20))
	t.INSERT(NewNode(6))
	t.INSERT(NewNode(7))
	t.INSERT(NewNode(3))
	fmt.Printf("%s\n", t)
	// Output:
	// [2 3 4 6 7 9 13 15 17 18 20]
}
