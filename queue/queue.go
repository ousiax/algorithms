package queue

import (
	"errors"
)

type Queue struct {
	q    [10]int // 使用数组实现一个最多容纳9个元素的队列
	tail int     // Q.tail 指向下一个新元素将要插入的位置
	head int     // Q.head 指向队头元素
}

// ENQUEUE(Q,x)
//   Q[Q.tail] = x
//   if Q.tail == Q.length
//     Q.tail = 1
//   else
//     Q.tail = Q.tail + 1
func (q *Queue) ENQUEUE(x int) error {
	if (q.tail == len(q.q)-1 && q.head == 0) || q.head == q.tail+1 {
		return errors.New("overflow")
	}

	q.q[q.tail] = x
	if q.tail == len(q.q)-1 {
		q.tail = 0
	} else {
		q.tail = q.tail + 1
	}
	return nil
}

// DEQUEUE(Q)
//   x = Q[Q.head]
//   if Q.head = Q.length
//     Q.head = 1
//   else
//     Q.head = Q.head + 1
//   return x
func (q *Queue) DEQUEUE() (int, error) {
	if q.head == q.tail {
		return 0, errors.New("underflow")
	}

	x := q.q[q.head]
	if q.head == len(q.q)-1 {
		q.head = 0
	} else {
		q.head = q.head + 1
	}
	return x, nil
}
