package bst

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	q := &nodeq{}

	if x := q.Deq(); x != nil {
		t.Errorf("q.Deq() return %v, want nil", x)
	}

	q.Enq(&Node{})

	if x := q.Deq(); x == nil {
		t.Errorf("q.Deq() return nil, want not nil")
	}

	if x := q.Deq(); x != nil {
		t.Errorf("q.Deq() return %v, want nil", x)
	}
}

func ExampleQ() {
	q := &nodeq{}
	q.Enq(&Node{key: 4})
	q.Enq(&Node{key: 2})
	q.Enq(&Node{key: 5})
	q.Enq(&Node{key: 1})
	q.Enq(&Node{key: 3})
	for x := q.Deq(); x != nil; x = q.Deq() {
		fmt.Printf("%d ", x.key)
	}
	// Output:
	// 4 2 5 1 3
}
