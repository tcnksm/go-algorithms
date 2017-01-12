package graph

import "testing"

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
}
