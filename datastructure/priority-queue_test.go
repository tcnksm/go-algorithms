package datastructure

import (
	"reflect"
	"testing"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func TestPriorityQueue(t *testing.T) {

	h := &intHeap{0}
	pq := NewPriorityQueue(h)

	pq.Enqueue(10)
	pq.Enqueue(5)
	pq.Enqueue(20)

	got := make([]int, 0, 4)
	for h.Len() > 0 {
		got = append(got, pq.Dequeue().(int))
	}

	want := []int{20, 10, 5, 0}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
