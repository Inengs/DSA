package sorting_algorithms

// Start with the unsorted slice a.
// Loop from the start to the end of the slice.
// For each pass, compare adjacent elements a[j] and a[j+1].
// If a[j] > a[j+1], swap them.
// After each full pass through the slice, the largest remaining element settles at the end.
// Repeat passes n-1 times or until a pass makes no swaps (early exit optimization).
// Result: sorted in ascending order.
// Time complexity: average/worst O(n²); best O(n) with early exit when already sorted.
// Space complexity: O(1) extra (in-place).
// Stable: yes.

func BubbleSort(array []int) {
	if len(array) < 2 {
		return
	}
	for i := 0; i < len(array)-1; i++ {
		swapped := false
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped { // if none is swapped, it means the array is already sorted
			return
		}
	}
}