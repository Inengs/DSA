package other_algorithms

// PRIM(G, start):
//     for each vertex v in G:
//         key[v] = ∞
//         parent[v] = NIL
//         inMST[v] = false

//     key[start] = 0

//     repeat V times:
//         u = vertex with minimum key not in MST
//         mark u as in MST

//         for each neighbor v of u:
//             if v not in MST and weight(u, v) < key[v]:
//                 key[v] = weight(u, v)
//                 parent[v] = u

//     return parent[]

import (
	"container/heap"
	"fmt"
)

// Edge represents a graph edge
type prims_Edge struct {
	to     int
	weight int
}

// Priority Queue Item
// type Item struct {
// 	vertex int
// 	key    int
// }

// Min Heap
type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Prim's Algorithm
func Prim(graph [][]prims_Edge, start int) {
	n := len(graph)
	inMST := make([]bool, n)
	key := make([]int, n)
	parent := make([]int, n)

	const INF = int(1e9)

	for i := 0; i < n; i++ {
		key[i] = INF
		parent[i] = -1
	}

	key[start] = 0
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, Item{node: start, priority: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(Item).node

		if inMST[u] {
			continue
		}

		inMST[u] = true

		for _, edge := range graph[u] {
			v := edge.to
			w := edge.weight

			if !inMST[v] && w < key[v] {
				key[v] = w
				parent[v] = u
				heap.Push(pq, Item{node: v, priority: w})
			}
		}
	}

	fmt.Println("Edges in the Minimum Spanning Tree:")
	for v := 1; v < n; v++ {
		fmt.Printf("%d -- %d\n", parent[v], v)
	}
}