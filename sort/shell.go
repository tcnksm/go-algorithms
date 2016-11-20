package sort

// Shell sort is improvement of Insert sort
func Shell(a []int) {
	for h := len(a) / 2; h > 0; h /= 2 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h && a[j-h] > a[j]; j -= h {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}
