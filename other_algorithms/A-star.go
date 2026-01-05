package other_algorithms

import (
	"container/heap"
)

type Node struct {
	id       int
	f, g, h  int
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func aStar(graph map[int]map[int]int, start, goal int, heuristic map[int]int) int {
	open := &PQ{}
	heap.Push(open, &Node{id: start, g: 0, h: heuristic[start], f: heuristic[start]})

	visited := make(map[int]bool)

	for open.Len() > 0 {
		current := heap.Pop(open).(*Node)

		if current.id == goal {
			return current.g
		}

		visited[current.id] = true

		for neighbor, cost := range graph[current.id] {
			if visited[neighbor] {
				continue
			}
			g := current.g + cost
			h := heuristic[neighbor]
			heap.Push(open, &Node{id: neighbor, g: g, h: h, f: g + h})
		}
	}
	return -1
}

