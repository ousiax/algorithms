package main

import (
	"fmt"
	"strings"

	l "./linkedlist/list"
)

func main() {
	t := HASH_TABLE{}
	t.INSERT(&l.Node{Key: 5})
	t.INSERT(&l.Node{Key: 28})
	t.INSERT(&l.Node{Key: 19})
	t.INSERT(&l.Node{Key: 15})
	t.INSERT(&l.Node{Key: 20})
	t.INSERT(&l.Node{Key: 33})
	t.INSERT(&l.Node{Key: 12})
	t.INSERT(&l.Node{Key: 17})
	t.INSERT(&l.Node{Key: 10})

	fmt.Printf("%s\n", t.String())

	// Output:
	// 00:
	// 01: 10 <=> 19 <=> 28
	// 02: 20
	// 03: 12
	// 04:
	// 05: 5
	// 06: 33 <=> 15
	// 07:
	// 08: 17

}

const slots_size = 9

type HASH_TABLE struct {
	slots [slots_size]l.LIST
}

func (h *HASH_TABLE) String() string {
	b := strings.Builder{}
	for s, l := range h.slots {
		b.WriteString(fmt.Sprintf("%02d: %s\n", s, &l))
	}
	return b.String()
}

func (h *HASH_TABLE) hash(key int) int {
	return key % slots_size
}

func (h *HASH_TABLE) INSERT(x *l.Node) {
	s := h.hash(x.Key)
	h.slots[s].INSERT(x)
}

func (h *HASH_TABLE) SEARCH(k int) *l.Node {
	s := h.hash(k)
	return h.slots[s].SEARCH(k)
}

func (h *HASH_TABLE) DELETE(x *l.Node) {
	s := h.hash(x.Key)
	h.slots[s].DELETE(x)
}
