package graph

import (
	"reflect"
	"testing"
)

func TestDirectedGraph(t *testing.T) {
	directedGraph := New(10, true)

	directedGraph.AddEdge(0, 1, 1)
	if !directedGraph.IsEdge(0, 1) {
		t.Fatalf("expect 0 -> 1 to be edge")
	}

	if directedGraph.IsEdge(1, 0) {
		t.Fatalf("expect 1 -> 0 not to be edge")
	}

	directedGraph.RemoveEdge(0, 1)
	if directedGraph.IsEdge(0, 1) {
		t.Fatalf("expect 0 -> 1 not to be edge (should be removed)")
	}

	directedGraph.AddEdge(0, 1, 1)
	directedGraph.AddEdge(0, 2, 1)
	directedGraph.AddEdge(0, 3, 1)
	if got, want := directedGraph.Neighbours(0), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Neighbours %#v, want %#v", got, want)
	}

}

func TestUndirectedGraph(t *testing.T) {
	undirectedGraph := New(10, false)

	undirectedGraph.AddEdge(0, 1, 1)
	if !undirectedGraph.IsEdge(0, 1) || !undirectedGraph.IsEdge(1, 0) {
		t.Fatalf("expect 1 - 0 to be edge")
	}

	undirectedGraph.RemoveEdge(0, 1)
	if undirectedGraph.IsEdge(0, 1) || undirectedGraph.IsEdge(1, 0) {
		t.Fatalf("expect 0 - 1 not to be edge")
	}

	undirectedGraph.AddEdge(0, 1, 1)
	undirectedGraph.AddEdge(0, 2, 1)
	undirectedGraph.AddEdge(0, 3, 1)
	if got, want := undirectedGraph.Neighbours(0), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Neighbours %#v, want %#v", got, want)
	}

}
