package graph

import "github.com/tcnksm/go-algorithms/datastructure"

func BreadthFirstSearch(graph *Graph, s int, fn func(i int)) error {
	visited := make([]int, graph.n)
	queue := &datastructure.Queue{}
	queue.Enqueue(s)

	for {
		v, err := queue.Dequeue()
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
			queue.Enqueue(j)
		}

	}
	return nil
}

func DepthFirstSearch(graph *Graph, s int, fn func(i int)) error {
	visited := make([]int, graph.n)
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
