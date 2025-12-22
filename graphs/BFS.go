package graphs

import "fmt"

// STEPS
// Start at a node
// This is your starting point (e.g., node 0).
// Mark it as visited
// This prevents coming back to it again and looping forever.
// Put it in a queue
// A queue follows First In, First Out (FIFO).
// This helps BFS visit nodes in the correct order.
// Take the first node from the queue
// This is the current node you are visiting.
// Visit all its immediate neighbors
// For each neighbor:
// If it has not been visited:
// Mark it as visited
// Add it to the queue
// Repeat
// Keep removing nodes from the queue and visiting their neighbors
// Continue until the queue is empty

func BFS(graph map[int][]int, start int) { // takes in a graph and the first value of the graph
	visited := make(map[int]bool) // mark visited columns
	queue := []int{start} // start a queue with the first value in the graph

	visited[start] = true // means i have visited the first value

	for len(queue) > 0 {
		node := queue[0] // take the first value
		queue = queue[1:] // all values from second to last

		fmt.Print(node, " ") // print out the node

		// loop through the graph and everyone that is visited mark it as visited
		for _, neighbour := range graph[node] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
}