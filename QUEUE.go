package main

import (
	"fmt"

	. "./queue"
)

func main() {
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
