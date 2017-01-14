package graph

import (
	"container/heap"

	"github.com/tcnksm/go-algorithms/datastructure"
)

const (
	inf int = 1 << 62
)

type vertex struct {
	id   int
	dist *int
}

type priorityQueue []vertex

func (q priorityQueue) Len() int           { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return *q[i].dist < *q[j].dist }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *priorityQueue) Push(v interface{}) {
	*q = append(*q, v.(vertex))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	v := old[n-1]
	*q = old[0 : n-1]
	return v
}

func Dijkstra(graph *datastructure.Graph, s int, dist []int) error {
	pq := make(priorityQueue, 0, graph.N)
	heap.Init(&pq)

	// Initialize dist
	for i := 0; i < graph.N; i++ {
		dist[i] = inf
		if i == s {
			dist[i] = 0
		}

		heap.Push(&pq, vertex{
			id:   i,
			dist: &dist[i],
		})
	}

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(vertex)
		for _, nextV := range graph.Neighbours(v.id) {
			newDist := dist[v.id] + graph.Edge(v.id, nextV)
			if newDist < dist[nextV] {
				dist[nextV] = newDist
			}
		}

		// Better to use Fix function
		// How to know index which is fixed
		heap.Init(&pq)
	}

	return nil
}
