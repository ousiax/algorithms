package main

import (
	"fmt"
	"math"
	"strings"

	l "./linkedlist/list"
)

func main() {
	t := HASH_TABLE{}
	t.INSERT(&l.Node{Key: 61})
	t.INSERT(&l.Node{Key: 62})
	t.INSERT(&l.Node{Key: 63})
	t.INSERT(&l.Node{Key: 64})
	t.INSERT(&l.Node{Key: 65})

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

const m = 1000

type HASH_TABLE struct {
	slots [m]l.LIST
}

func (h *HASH_TABLE) String() string {
	b := strings.Builder{}
	for s, l := range h.slots {
		ls := fmt.Sprintf("%s", &l)
		if ls != "" {
			b.WriteString(fmt.Sprintf("%03d: %s\n", s, ls))
		}
	}
	return b.String()
}

func (h *HASH_TABLE) hash(k int) int {
	// A = (sqrt(5) - 1) / 2
	// hash(k) = [m(kA mod 1)]  // round down
	return int(math.Floor(float64(m) * math.Mod(float64(k)*((math.Sqrt(5)-1)/2), 1)))
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
