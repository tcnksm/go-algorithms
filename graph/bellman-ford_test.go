package graph

import (
	"reflect"
	"testing"

	"github.com/tcnksm/go-algorithms/datastructure"
)

func testGraphWithNegative(t *testing.T) (*datastructure.Graph, []int) {
	directedGraph := datastructure.NewGraph(5, true)

	directedGraph.AddEdge(0, 4, 2)
	directedGraph.AddEdge(4, 3, 4)
	directedGraph.AddEdge(4, 1, 5)
	directedGraph.AddEdge(3, 2, 6)
	directedGraph.AddEdge(1, 3, -2)
	directedGraph.AddEdge(2, 1, -3)

	want := []int{0, 7, 11, 5, 2}

	return directedGraph, want
}

func TestBellmanFord(t *testing.T) {
	directedGraph, want := testGraphWithNegative(t)
	dist := make([]int, 5)
	if err := BellmanFord(directedGraph, 0, dist); err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(dist, want) {
		t.Fatalf("got %v, want %v", dist, want)
	}
}

func TestBellmanFord_negative_cycle(t *testing.T) {
	directedGraph := datastructure.NewGraph(5, true)

	directedGraph.AddEdge(0, 4, 2)
	directedGraph.AddEdge(4, 3, 4)
	directedGraph.AddEdge(4, 1, 5)
	directedGraph.AddEdge(3, 2, 6)
	directedGraph.AddEdge(1, 3, -2)
	directedGraph.AddEdge(2, 1, -5)

	dist := make([]int, 5)
	if err := BellmanFord(directedGraph, 0, dist); err == nil {
		t.Fatalf("expect to be failed: %v", dist)
	}
}
