package datastructure

import (
	"container/heap"
)

type PriorityQueue struct {
	h heap.Interface
}

func NewPriorityQueue(h heap.Interface) *PriorityQueue {
	heap.Init(h)
	return &PriorityQueue{
		h: h,
	}
}

func (q *PriorityQueue) Enqueue(v interface{}) {
	heap.Push(q.h, v)
}

func (q *PriorityQueue) Dequeue() interface{} {
	return heap.Pop(q.h)
}
