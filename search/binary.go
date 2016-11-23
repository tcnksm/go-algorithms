package search

// Binary search search sorted slice a contains x.
//
// Best:    O(1)
// Worst:   O(log n)
// Average: O(log n)
func Binary(a []int, x int) bool {

	left, right := 0, len(a)
	for left < right {
		index := left + (right-left)/2
		if a[index] == x {
			return true
		}

		if x < a[index] {
			right = index - 1
		} else {
			left = index + 1
		}
	}
	return false
}
