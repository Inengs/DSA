package other_algorithms

// KRUSKAL(G):
//     MST = empty set

//     sort all edges of G by increasing weight

//     for each vertex v in G:
//         makeSet(v)

//     for each edge (u, v) in sorted edges:
//         if find(u) != find(v):
//             add (u, v) to MST
//             union(u, v)

//     return MST

import (
	"sort"
)

// Edge represents a weighted edge
// type Edge struct {
// 	u, v, weight int
// }

// Disjoint Set (Union-Find)
type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}
	return ds
}

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX == rootY {
		return
	}

	if ds.rank[rootX] < ds.rank[rootY] {
		ds.parent[rootX] = rootY
	} else if ds.rank[rootX] > ds.rank[rootY] {
		ds.parent[rootY] = rootX
	} else {
		ds.parent[rootY] = rootX
		ds.rank[rootX]++
	}
}

// Kruskal's Algorithm
func Kruskal(vertices int, edges []Edge) []Edge {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	ds := NewDisjointSet(vertices)
	mst := []Edge{}

	for _, edge := range edges {
		if ds.Find(edge.u) != ds.Find(edge.v) {
			mst = append(mst, edge)
			ds.Union(edge.u, edge.v)
		}
	}

	return mst
}
