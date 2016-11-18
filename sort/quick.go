package sort

func Quick(a []int) {
	quick(a, 0, len(a)-1)
}

func quick(a []int, left, right int) {
	if left >= right {
		return
	}

	pi := partition(a, left, right)
	quick(a, left, pi-1)
	quick(a, pi+1, right)
}

func partition(a []int, left, right int) int {
	p := left
	a[p], a[right] = a[right], a[p]

	store := left
	for i := left; i < right; i++ {
		if a[i] <= a[right] {
			a[store], a[i] = a[i], a[store]
			store++
		}
	}

	a[right], a[store] = a[store], a[right]
	return store
}

// QuickSimple is simple implementation of quick sort
// but it's slow and needs a lot of allocation.
func QuickSimple(a []int) {
	quickSimple(a)
}

func quickSimple(a []int) []int {
	if len(a) < 1 {
		return a
	}

	pivot := a[(len(a)-1)/2]
	left := make([]int, 0, len(a)/2)
	right := make([]int, 0, len(a)/2)
	for i := 1; i < len(a); i++ {
		switch {
		case a[i] <= pivot:
			left = append(left, a[i])
		case a[i] > pivot:
			right = append(right, a[i])
		}
	}

	left = quickSimple(left)
	right = quickSimple(right)

	left = append(left, pivot)
	left = append(left, right...)
	return left
}
