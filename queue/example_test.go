package queue

import (
	"fmt"
)

func ExampleQueue() {
	var q = new(Queue)
	q.ENQUEUE(4)
	q.ENQUEUE(1)
	q.ENQUEUE(3)
	q.DEQUEUE()
	q.ENQUEUE(8)
	q.DEQUEUE()
	for {
		if x, err := q.DEQUEUE(); err == nil {
			fmt.Printf("%d ", x)
		} else {
			break
		}
	}
	// Output:
	// 3 8
}

func ExampleENDEQUEUE() {
	q := new(Queue)
	i := 11
	for {
		if err := q.ENQUEUE(i); err != nil {
			break
		}
		i += 11
	}

	for {
		if x, err := q.DEQUEUE(); err == nil {
			fmt.Printf("%d ", x)
		} else {
			break
		}
	}

	// Output:
	// 11 22 33 44 55 66 77 88 99
}
