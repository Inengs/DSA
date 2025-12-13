package sorting_algorithms

// SELECTION SORT IMPLEMENTATION
// 1. Loop through the array
// 2. Assume current index is the minimum
// 3. Find the smallest element in the unsorted part
// 4. Swap it with the current index
// 5. Move to the next position

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
