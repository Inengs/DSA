package sorting_algorithms

// Pick the subarray you want to sort, defined by low and high (inclusive).
// Choose a pivot value — often a[(low+high)/2] (middle element).
// Set left = low and right = high.
// Enter a loop that runs until left > right (or pointers cross).
// Move left rightwards: while a[left] < pivot do left++.
// Move right leftwards: while a[right] > pivot do right--.
// When both pointers stopped, check: if left >= right, break and return right as partition index.
// If left < right, swap a[left] and a[right].
// After swap, advance left++ and right-- and continue the loop.
// When loop ends, right is the split point: elements low..right ≤ (roughly) pivot, right+1..high ≥ pivot.
// Recursively quicksort the left subarray low..right.
// Recursively quicksort the right subarray right+1..high.
// Base case: if low >= high, return (0 or 1 element).
// Practical tweaks: use middle pivot or randomize pivot; recurse on smaller side first to limit stack depth; switch to insertion sort for tiny subarrays.
// Note: Hoare’s partition is in-place, usually fewer swaps than Lomuto, and handles duplicates well — but the partition index is not the pivot’s final position.

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, low, high int) {
	if low >= high {
		return
	}

	p := partition(a, low, high)

	// Hoare partition returns a split point that is not the final pivot index,
	// but both sides are valid to recurse on.
	quickSort(a, low, p)
	quickSort(a, p+1, high)
}

func partition(a []int, low, high int) int {
	pivot := a[(low+high)/2] // choose middle element as pivot

	itemFromLeft := low - 1
	itemFromRight := high + 1

	for {
		// move left pointer rightward until finding >= pivot
		for {
			itemFromLeft++
			if a[itemFromLeft] >= pivot {
				break
			}
		}

		// move right pointer leftward until finding <= pivot
		for {
			itemFromRight--
			if a[itemFromRight] <= pivot {
				break
			}
		}

		// if pointers cross, partition is done
		if itemFromLeft >= itemFromRight {
			return itemFromRight
		}

		// otherwise swap the two items
		a[itemFromLeft], a[itemFromRight] = a[itemFromRight], a[itemFromLeft]
	}
}