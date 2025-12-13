package sorting_algorithms

// MERGE SORT IMPLEMENTATION
// 1. Base case: array with 0 or 1 element is already sorted
// 2. Split the array into two halves
// 3. Recursively sort both halves
// 4. Merge the two sorted halves into one sorted array

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

// Helper function to merge two sorted slices
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
