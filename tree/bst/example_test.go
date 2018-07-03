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

	fmt.Printf("PREORDER: ")
	t.PREORDER_WALK(func(x *Node) { fmt.Printf("%d,", x.Key()) })
	fmt.Println()

	fmt.Printf("INORDER: ")
	t.INORDER_WALK(func(x *Node) { fmt.Printf("%d,", x.Key()) })
	fmt.Println()

	fmt.Printf("POSTORDER: ")
	t.POSTORDER_WALK(func(x *Node) { fmt.Printf("%d,", x.Key()) })
	fmt.Println()

	fmt.Printf("LEVELORDER: ")
	t.LEVELORDER_WALK(func(x *Node) { fmt.Printf("%d,", x.Key()) })
	fmt.Println()
	// Output:
	// [2 3 4 6 7 9 13 15 17 18 20]
	// PREORDER: 15,4,2,3,13,9,6,7,18,17,20,
	// INORDER: 2,3,4,6,7,9,13,15,17,18,20,
	// POSTORDER: 15,4,2,3,13,9,6,7,18,17,20,
	// LEVELORDER: 15,4,18,2,13,17,20,3,9,6,7,
}
