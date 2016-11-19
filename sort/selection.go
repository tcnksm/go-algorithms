package sort

func Selection(a []int) {
	for i := 0; i < len(a); i++ {
		j := selectMin(a, i)
		if j != i {
			a[i], a[j] = a[j], a[i]
		}
	}
}

func selectMin(a []int, left int) int {
	index := left
	for i := left; i < len(a); i++ {
		if a[i] < a[index] {
			index = i
		}
	}

	return index
}
