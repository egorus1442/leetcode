package structs

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	data     []int
	capacity int
	head     int
	tail     int
	size     int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]int, capacity),
		capacity: capacity,
		head:     0,
		tail:     -1,
		size:     0,
	}
}

func (q *ArrayQueue) Enqueue(value int) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	q.tail = (q.tail + 1) % q.capacity
	q.data[q.tail] = value
	q.size++
	return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}
	value := q.head
	q.head = (q.head + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) PrintQueue() {
	fmt.Println(q.data)
}
