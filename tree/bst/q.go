package bst

type nodeq struct {
	head *node
}

type node struct {
	next *node
	v    *Node
}

func (nq *nodeq) Enq(x *Node) {
	if x != nil {
		if nq.head == nil {
			nq.head = &node{v: x}
		} else {
			for t := nq.head; ; t = t.next {
				if t.next == nil {
					t.next = &node{v: x}
					break
				}
			}
		}
	}
}

func (nq *nodeq) Deq() *Node {
	if nq.head == nil {
		return nil
	}

	x := nq.head.v
	nq.head = nq.head.next
	return x
}
