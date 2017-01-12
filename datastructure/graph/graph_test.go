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

func TestDirectedGraph_DFS(t *testing.T) {
	directedGraph := New(10)

	// Graph to construct below,
	//
	//            0
	//       /    |   \
	//     1      2      3
	//    / \    / \    /
	//   4   5  6   7  8
	//  /
	// 9
	directedGraph.AddEdge(0, 1, 1)
	directedGraph.AddEdge(0, 2, 1)
	directedGraph.AddEdge(0, 3, 1)
	directedGraph.AddEdge(1, 4, 1)
	directedGraph.AddEdge(1, 5, 1)
	directedGraph.AddEdge(2, 6, 1)
	directedGraph.AddEdge(2, 7, 1)
	directedGraph.AddEdge(3, 8, 1)
	directedGraph.AddEdge(4, 9, 1)

	want := []int{0, 3, 8, 2, 7, 6, 1, 5, 4, 9}
	got := make([]int, 0, 10)
	directedGraph.DepthFirstSearch(0, func(i int) {
		got = append(got, i)
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

}
