package structs

import (
	"errors"
)

type Node struct {
	value int
	next  *Node
}

type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *LinkedListQueue) EnLinkedListQueue(value int) {
	newNode := &Node{value: value, next: nil}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *LinkedListQueue) DeLinkedListQueue() (int, error) {
	if q.head == nil {
		return 0, errors.New("queue is empty")
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return value, nil
}

func (q *LinkedListQueue) SizeLinkedListQueue() int {
	return q.size
}
