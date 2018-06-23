package list

import (
	"fmt"
)

func ExampleLIST() {
	L := new(LIST)
	x := &Node{Key: 1}
	L.INSERT(x)
	L.INSERT(&Node{Key: 4})
	L.INSERT(&Node{Key: 16})
	L.INSERT(&Node{Key: 9})
	fmt.Printf("%v\n", L)
	L.INSERT(&Node{Key: 25})
	fmt.Printf("%v\n", L)
	L.DELETE(x)
	fmt.Printf("%v\n", L)

	// Output:
	// 9 <=> 16 <=> 4 <=> 1
	// 25 <=> 9 <=> 16 <=> 4 <=> 1
	// 25 <=> 9 <=> 16 <=> 4
}
