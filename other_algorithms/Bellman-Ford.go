package other_algorithms

import (
	"math"
)

type Edge struct {
	u, v, w int
}

func bellmanFord(edges []Edge, V, start int) ([]int, bool) {
	dist := make([]int, V)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	for i := 0; i < V-1; i++ {
		for _, e := range edges {
			if dist[e.u] != math.MaxInt && dist[e.u]+e.w < dist[e.v] {
				dist[e.v] = dist[e.u] + e.w
			}
		}
	}

	// Check negative cycle
	for _, e := range edges {
		if dist[e.u] != math.MaxInt && dist[e.u]+e.w < dist[e.v] {
			return dist, true
		}
	}
	return dist, false
}

