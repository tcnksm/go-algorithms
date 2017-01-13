package graph

import (
	"reflect"
	"testing"
)

func testGraph(t *testing.T) *Graph {
	directedGraph := New(10)

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

	return directedGraph
}

func TestDirectedGraph_DFS(t *testing.T) {
	graph := testGraph(t)
	want := []int{0, 3, 8, 2, 7, 6, 1, 5, 4, 9}

	got := make([]int, 0, 10)
	err := DepthFirstSearch(graph, 0, func(i int) {
		got = append(got, i)
	})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

}

func TestDirectedGraph_BFS(t *testing.T) {
	graph := testGraph(t)
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := make([]int, 0, 10)
	err := BreadthFirstSearch(graph, 0, func(i int) {
		got = append(got, i)
	})

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
