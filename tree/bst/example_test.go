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

func ExampleBST_RECURISON() {
	t := &Tree{}
	t.INSERT_RECURSION(NewNode(15))
	t.INSERT_RECURSION(NewNode(4))
	t.INSERT_RECURSION(NewNode(2))
	t.INSERT_RECURSION(NewNode(13))
	t.INSERT_RECURSION(NewNode(9))
	t.INSERT_RECURSION(NewNode(18))
	t.INSERT_RECURSION(NewNode(17))
	t.INSERT_RECURSION(NewNode(20))
	t.INSERT_RECURSION(NewNode(6))
	t.INSERT_RECURSION(NewNode(7))
	t.INSERT_RECURSION(NewNode(3))

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
