package main

type MyCircularQueue struct {
	data     []int
	capacity int
	head     int
	tail     int
	size     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:     make([]int, k),
		capacity: k,
		head:     0,
		tail:     -1,
		size:     0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == this.capacity {
		return false
	}
	this.tail = (this.tail + 1) % this.capacity
	this.data[this.tail] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.size == this.capacity {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
