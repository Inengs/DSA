package graphs

import "fmt"

func AdjListGraph() {
	graph := make(map[int][]int) // create a map with key of int and values of array of int

	graph[0] = []int{1, 2} // 1 and 2 are neighbouring nodes
	graph[1] = []int{2}    // 2 is the only neighbouring node
	graph[2] = []int{0, 3} // 0 and 3 are the neighbouring nodes
	graph[3] = []int{} // it has no neighbouring nodes

	fmt.Println("Adjacency list")

	for node, neighbours := range graph {
		fmt.Println(node, neighbours)
	}
}

func AdjMatrixGraph() {
	// Graph represented as adjacency matrix
	graph := [][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{1, 0, 0, 1},
		{0, 0, 0, 0},
	}

	fmt.Println("Adjacency Matrix:")
	for _, row := range graph {
		fmt.Println(row) // print out each row
	}
}