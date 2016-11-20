package sort

// Counting can be used only if elements of a is
// 0 <= i < k and k should be much smaller than len(a).
func Counting(a []int, k int) {
	b := make(map[int]int, k)
	for i := 0; i < len(a); i++ {
		b[a[i]]++
	}

	index := 0
	for i := 0; i <= k; i++ {
		for b[i] > 0 {
			a[index] = i
			b[i]--
			index++
		}
	}
}
