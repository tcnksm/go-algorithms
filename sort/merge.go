package sort

func Merge(a []int) {
	merge(a, 0, len(a)-1)
}

func merge(a []int, start, end int) {
	if start >= end {
		return
	}

	m := (start + end) / 2
	merge(a, start, m)
	merge(a, m+1, end)

	// It's difficult to merge 2 array in-place
	var index int
	b := make([]int, end-start+1)

	i := start
	j := m + 1
	for {
		if i > m || j > end {
			break
		}

		if a[i] <= a[j] {
			b[index] = a[i]
			index++
			i++
		} else {
			b[index] = a[j]
			index++
			j++
		}
	}

	for i <= m {
		b[index] = a[i]
		index++
		i++
	}

	for j <= end {
		b[index] = a[j]
		index++
		j++
	}

	n := 0
	for k := start; k <= end; k++ {
		a[k] = b[n]
		n++
	}
}
