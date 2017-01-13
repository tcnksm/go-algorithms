package graph

import "fmt"

type Graph struct {
	n        int // number of vertex
	edges    [][]int
	directed bool
}

func New(n int) *Graph {
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		edges[i] = make([]int, n)
	}

	return &Graph{
		n:        n,
		edges:    edges,
		directed: true,
	}
}

func (g *Graph) AddEdge(i, j int, w int) error {
	if i == j {
		return fmt.Errorf("self-loop")
	}

	if w < 1 {
		return fmt.Errorf("weight must be more than 1")
	}

	// Not allow to add edge twice
	if v := g.edges[i][j]; v != 0 {
		return fmt.Errorf("edge already exists")
	}

	g.edges[i][j] = w
	return nil
}

func (g *Graph) RemoveEdge(i, j int) error {
	if v := g.edges[i][j]; v == 0 {
		return fmt.Errorf("edge doesn't exist")
	}

	g.edges[i][j] = 0
	return nil
}

func (g *Graph) IsEdge(i, j int) bool {
	if v := g.edges[i][j]; v != 0 {
		return true
	}
	return false
}

func (g *Graph) Neighbours(i int) []int {
	neighbours := make([]int, 0, g.n)
	for j := 0; j < g.n; j++ {
		v := g.edges[i][j]
		if v > 0 {
			neighbours = append(neighbours, j)
		}
	}
	return neighbours
}
