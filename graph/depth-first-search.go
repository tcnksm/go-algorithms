package graph

import "github.com/tcnksm/go-algorithms/datastructure"

// O(V+E)
//
// V is number of vertics, E is number of edges
func DepthFirstSearch(graph *datastructure.Graph, s int, fn func(i int)) error {
	visited := make([]int, graph.N)
	stack := &datastructure.Stack{}
	stack.Push(s)

	for {
		v, err := stack.Pop()
		if err != nil {
			break
		}

		i := v.(int)
		if visited[i] == 1 {
			continue
		}

		visited[i] = 1
		fn(i)

		for _, j := range graph.Neighbours(i) {
			stack.Push(j)
		}
	}

	return nil
}
