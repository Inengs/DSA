package graphs

import "fmt"

// STEPS
// Start at a node
// This is your starting point.
// Mark it as visited
// So you don’t visit it again.
// Move to one unvisited neighbor
// Pick any neighbor that hasn’t been visited.
// Repeat the process for that neighbor
// From that neighbor, move to one of its unvisited neighbors.
// Go as deep as possible
// Keep moving forward until:
// There are no unvisited neighbors left
// Backtrack
// Go back to the previous node
// Try another unvisited neighbor
// Continue
// Repeat until all reachable nodes are visited

func DFS(graph map[int][]int, node int, visited map[int]bool) { // takes in a graph, a node and the a visited map
	visited[node] = true // give the value to the first node key visited
	fmt.Print(node, "") // print out that node

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			DFS(graph, neighbor, visited)
		}
	}
}