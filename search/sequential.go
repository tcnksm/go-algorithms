package search

// Sequential search
//
// Best:    O(1)
// Worst:   O(n)
// Average: O(n)
func Sequential(a []int, x int) bool {
	for _, e := range a {
		if e == x {
			return true
		}
	}
	return false
}
