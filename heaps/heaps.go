package heaps

type MaxHeap struct {
	data []int
}

// Get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get left child index
func leftChild(i int) int {
	return 2*i + 1
}

// Get right child index
func rightChild(i int) int {
	return 2*i + 2
}

// Insert element into heap
func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// Heapify up (after insertion)
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 && h.data[parent(index)] < h.data[index] {
		h.data[parent(index)], h.data[index] =
			h.data[index], h.data[parent(index)]
		index = parent(index)
	}
}

// Extract max element (root)
func (h *MaxHeap) ExtractMax() int {
	if len(h.data) == 0 {
		return -1
	}

	max := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.heapifyDown(0)
	return max
}

// Heapify down (after deletion)
func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.data) - 1

	for {
		left := leftChild(index)
		right := rightChild(index)
		largest := index

		if left <= lastIndex && h.data[left] > h.data[largest] {
			largest = left
		}

		if right <= lastIndex && h.data[right] > h.data[largest] {
			largest = right
		}

		if largest == index {
			break
		}

		h.data[index], h.data[largest] =
			h.data[largest], h.data[index]

		index = largest
	}
}

