package search

// Sequential search
//
// Best:    O(1)
// Worst:   O(n)
// Average: O(n)
func Sequential(a []int, x int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return true
		}
	}
	return false
}
