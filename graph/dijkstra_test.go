package graph

import (
	"reflect"
	"testing"

	"github.com/tcnksm/go-algorithms/datastructure"
)

func TestDijkstra(t *testing.T) {
	directedGraph := datastructure.NewGraph(6, true)

	directedGraph.AddEdge(0, 1, 6)
	directedGraph.AddEdge(0, 2, 8)
	directedGraph.AddEdge(0, 3, 18)
	directedGraph.AddEdge(1, 4, 11)
	directedGraph.AddEdge(4, 5, 3)
	directedGraph.AddEdge(5, 3, 4)
	directedGraph.AddEdge(5, 2, 7)
	directedGraph.AddEdge(2, 3, 9)

	// dist[5] mean shortest distance (weight)
	// from vertex 0 to vertex 5
	want := []int{0, 6, 8, 17, 17, 20}

	dist := make([]int, 6)
	if err := Dijkstra(directedGraph, 0, dist); err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(dist, want) {
		t.Fatalf("got %v, want %v", dist, want)
	}
}
