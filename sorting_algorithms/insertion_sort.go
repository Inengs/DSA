package sorting_algorithms

// INSERTION SORT IMPLEMENTATION
// 1. Start from index 1
// 2. Store current value as key
// 3. Compare key with elements to the left
// 4. Shift larger elements to the right
// 5. Insert key in the correct position

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}
