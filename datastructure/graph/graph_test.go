package graph

import (
	"reflect"
	"testing"
)

func TestDirectedGraph(t *testing.T) {
	directedGraph := New(10)

	directedGraph.AddEdge(0, 1, 1)
	if !directedGraph.IsEdge(0, 1) {
		t.Fatalf("expect 1 -> 1 to be edge")
	}

	directedGraph.RemoveEdge(0, 1)
	if directedGraph.IsEdge(0, 1) {
		t.Fatalf("expect 1 -> 1 not to be edge")
	}

	directedGraph.AddEdge(0, 1, 1)
	directedGraph.AddEdge(0, 2, 1)
	directedGraph.AddEdge(0, 3, 1)
	if got, want := directedGraph.Neighbours(0), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Neighbours %#v, want %#v", got, want)
	}

}
