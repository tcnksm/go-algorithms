package graph

import (
	"fmt"

	"github.com/tcnksm/go-algorithms/datastructure"
)

func BellmanFord(graph *datastructure.Graph, s int, dist []int) error {
	for i := 0; i < len(dist); i++ {
		if i == s {
			dist[i] = 0
			continue
		}
		dist[i] = inf
	}

	// We need, at least, n-1 times inspection
	for i := 0; i <= graph.N; i++ {

		for u := 0; u < graph.N; u++ {
			for _, v := range graph.Neighbours(u) {
				newDist := dist[u] + graph.Edge(u, v)
				if newDist < dist[v] {
					if i == graph.N {
						return fmt.Errorf("graph has negative cycle")
					}

					dist[v] = newDist
				}
			}
		}
	}
	return nil
}
