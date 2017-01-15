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

// O((V+E)logV)
func Dijkstra(graph *datastructure.Graph, s int, dist []int) error {
	pq := make(priorityQueue, 0, graph.N)
	heap.Init(&pq)

	// Initialize dist, O(V*logV)
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

	// O(E*logV)
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(vertex).id
		for _, v := range graph.Neighbours(u) {
			newDist := dist[u] + graph.Edge(u, v)
			if newDist < dist[v] {
				dist[v] = newDist
			}
		}

		// Better to use Fix function
		// How to know index which is fixed
		heap.Init(&pq)
	}

	return nil
}

// Without heap
func DijkstraDG(graph *datastructure.Graph, s int, dist []int) error {
	visited := make([]bool, graph.N, graph.N)
	for i := 0; i < graph.N; i++ {
		if i == s {
			dist[i] = 0
			continue
		}
		dist[i] = inf
	}

	for {

		// Find V(u) which is not visited
		u := -1
		for i := 0; i < graph.N; i++ {
			if !visited[i] && dist[i] < inf {
				u = i
			}
		}

		if u == -1 {
			break
		}

		visited[u] = true
		for _, v := range graph.Neighbours(u) {
			newDist := dist[u] + graph.Edge(u, v)
			if newDist < dist[v] {
				dist[v] = newDist
			}
		}
	}

	return nil
}
