package queue

import "errors"

type node struct {
	number int
	next   *node
	prev   *node
}

type Queue struct {
	head *node
	tail *node
}

func (q *Queue) Enqueue(number int) {
	if q.head == nil && q.tail == nil {
		q.head = &node{number: number}
		q.tail = q.head
	} else {
		tmp := &node{number: number, next: q.head}
		q.head.prev = tmp
		q.head = tmp
	}
}

func (q *Queue) Dequeue() (int, error) {
	if q.head == nil && q.tail == nil {
		return 0, errors.New("Nothing to dequeue")
	}

	var result int
	if q.head == q.tail {
		ret := q.head
		q.head, q.tail = nil, nil
		result = ret.number
	} else {
		ret := q.tail
		ret.prev.next = nil
		q.tail = ret.prev
		result = ret.number
	}
	return result, nil
}
