package search_algorithms

import "fmt"

// When to use: sorted slices only. Significantly faster on large sorted arrays.

// Idea (iterative): keep low and high indices; compute mid = low + (high-low)/2; compare target to arr[mid] and move low or high until found or low > high.

// Ensure the slice is sorted (precondition).

// Initialize low := 0, high := len(arr)-1.

// While low <= high, compute mid safely as low + (high-low)/2.

// If arr[mid] == target → return mid.

// If target < arr[mid] → high = mid - 1.

// Else → low = mid + 1.

// If loop ends, return -1.

func BinarySearchIntIterative(arr []int, target int) int {
	low := 0
	high := len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if (arr[mid] == target) {
			return mid
		} 
		if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1 
}

func ShowResultBinaryIterative() {
	sorted := []int{1, 3, 4, 7, 9, 12}
	fmt.Println("Sorted:", sorted)
	fmt.Println("Search 7", BinarySearchIntIterative(sorted, 7))
}

// Complexity:
// Time: O(log n)
// Space: O(1)

// Binary search Recursive method
// Start with a sorted array and a target value.
// Use a wrapper to call the recursive function with low = 0 and high = len(arr)-1.
// If low > high, return -1 (interval empty).
// Compute mid = low + (high-low)/2.
// If arr[mid] == target, return mid (found).
// If target < arr[mid], recurse on left half: [low, mid-1].
// If target > arr[mid], recurse on right half: [mid+1, high].
// Each recursive call reduces the interval size.
// Recursion ends when the value is found or interval becomes empty.

func BinarySearchIntRecursive(arr []int, target int) int {
	return binaryRec(arr, target, 0, len(arr)-1)
}

func binaryRec(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid
	}
	if target < arr[mid] {
		return binaryRec(arr, target, low, mid-1)
	}
	return binaryRec(arr, target, mid+1, high)
}


func ShowResultBinaryRecursive() {
	sorted := []int{2, 4, 4, 6, 8}
	fmt.Println("Search 4 ->", BinarySearchIntRecursive(sorted, 4))
}