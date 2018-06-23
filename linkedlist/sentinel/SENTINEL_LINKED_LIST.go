package sentinel

import (
	"strconv"
	"strings"
)

type list struct {
	null *Node
}

func NewLIST() *list {
	l := &list{null: &Node{}}
	l.null.Next = l.null
	l.null.Prev = l.null
	return l
}

func (L *list) head() *Node {
	return L.null.Next
}

func (L *list) tail() *Node {
	return L.null.Prev
}

func (L *list) String() string {
	b := strings.Builder{}
	x := L.null.Next
	for x != L.null {
		b.WriteString(strconv.Itoa(x.Key))
		if x.Next != L.null {
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
//   x = L.nil.next
//   while x != L.nil and x.key != k
//     x = x.next
func (L *list) SEARCH(k int) *Node {
	x := L.null.Next
	for x != L.null && x.Key != k {
		x = x.Next
	}
	return x
}

// LIST-INSERT(L, x)
//   x.next = L.nil.next
//   L.nil.next.prev = x
//   L.nil.next = x
//   x.prev = L.nil
func (L *list) INSERT(x *Node) {
	x.Next = L.null.Next
	L.null.Next.Prev = x
	L.null.Next = x
	x.Prev = L.null

}

// LIST-DELETE(L,x)
//   x.prev.next = x.next
//   x.next.prev = x.prev
func (L *list) DELETE(x *Node) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
}
