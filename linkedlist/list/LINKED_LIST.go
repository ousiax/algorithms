package list

import (
	"strconv"
	"strings"
)

type LIST struct {
	head *Node
}

func (L *LIST) String() string {
	b := strings.Builder{}
	x := L.head
	for x != nil {
		b.WriteString(strconv.Itoa(x.Key))
		if x.Next != nil {
			b.WriteString(" <=> ")
		}
		x = x.Next
	}
	return b.String()
}

type Node struct {
	Prev *Node
	Next *Node
	Key  int
}

// LIST-SEARCH(L,k)
//   x = L.head
//   while x != NIL and x.key != k
//     x = x.next
func (L *LIST) SEARCH(k int) *Node {
	x := L.head
	for x != nil && x.Key != k {
		x = x.Next
	}
	return x
}

// LIST-INSERT(L, x)
//   x.next = L.head
//   if L.head != NIL
//     L.head.prev = x
//   L.head = x
//   x.prev = NIL
func (L *LIST) INSERT(x *Node) {
	x.Next = L.head
	if L.head != nil {
		L.head.Prev = x
	}
	L.head = x
	x.Prev = nil
}

// LIST-DELETE(L,x)
//   if x.prev != NIL
//     x.prev.next = x.next
//   else
//     L.head = x.next
//   if x.next != NIL
//     x.next.prev = x
func (L *LIST) DELETE(x *Node) {
	if x.Prev != nil {
		x.Prev.Next = x.Next
	} else {
		L.head = x.Next
	}
	if x.Next != nil {
		x.Next.Prev = x.Prev
	}
}
