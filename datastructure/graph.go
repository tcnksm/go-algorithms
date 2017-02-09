package datastructure

import "fmt"

type Graph struct {
	N        int // number of verte
	edges    [][]int
	directed bool
}

const (
	testVallue = 100
)

func NewGraph(n int, directed bool) *Graph {
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		edges[i] = make([]int, n)
	}

	return &Graph{
		N:        n,
		edges:    edges,
		directed: directed,
	}
}

func (g *Graph) Edge(i, j int) int {
	return g.edges[i][j]
}

func (g *Graph) AddEdge(i, j int, w int) error {
	if i == j {
		return fmt.Errorf("self-loop")
	}

	if w == 0 {
		return fmt.Errorf("Weight must be non-zero value")
	}

	// Not allow to add edge twice
	if v := g.edges[i][j]; v != 0 {
		return fmt.Errorf("edge already exists")
	}

	g.edges[i][j] = w

	if !g.directed {
		g.edges[j][i] = w
	}

	return nil
}

func (g *Graph) RemoveEdge(i, j int) error {
	if v := g.edges[i][j]; v == 0 {
		return fmt.Errorf("edge doesn't exist")
	}

	g.edges[i][j] = 0

	if !g.directed {
		g.edges[j][i] = 0
	}
	return nil
}

func (g *Graph) IsEdge(i, j int) bool {
	if v := g.edges[i][j]; v != 0 {
		return true
	}
	return false
}

func (g *Graph) Neighbours(i int) []int {
	neighbours := make([]int, 0, g.N)
	for j := 0; j < g.N; j++ {
		v := g.edges[i][j]
		if v != 0 {
			neighbours = append(neighbours, j)
		}
	}
	return neighbours
}
