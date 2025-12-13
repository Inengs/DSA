package sorting_algorithms

// HEAP SORT IMPLEMENTATION
// 1. Build a max heap from the array
// 2. Swap the root with the last element
// 3. Reduce heap size
// 4. Heapify the root
// 5. Repeat until sorted

func HeapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		Heapify(arr, n, i)
	}

	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		Heapify(arr, i, 0)
	}
}

// Heapify subtree rooted at index i
func Heapify(arr []int, heapSize int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, heapSize, largest)
	}
}
