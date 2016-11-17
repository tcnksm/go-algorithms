package sort

// Insert sorts https://en.wikipedia.org/wiki/Insertion_sort
func Insert(a []int) {
	for i := 0; i < len(a); i++ {
		// Get element one by one and compare with elements
		// which are already sorted.
		j := i
		for j > 0 && a[j-1] > a[j] {
			a[j-1], a[j] = a[j], a[j-1]
			j--
		}
	}
}
