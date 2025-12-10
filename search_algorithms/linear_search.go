package search_algorithms

import "fmt"

// scan each element and compare, return the index if found, else return -1

// Accept a slice and the target value.

// Loop from 0 to len(slice)-1.

// If current element == target, return index.

// If loop finishes, return -1.

func LinearSearchInt(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func ShowResult() {
	arr := []int{7,3,9,1,4}
	fmt.Println("Array:", arr)
	fmt.Println("Search 9", LinearSearchInt(arr, 9))
}

// Time: O(n) (must potentially check every element)

// Space: O(1)