package sort

func Heap(a []int) {
	buildHeap(a)
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, 0, i)
	}

}

func buildHeap(a []int) {
	for i := (len(a))/2 - 1; i >= 0; i-- {
		heapify(a, i, len(a))
	}
}

func heapify(a []int, index, max int) {
	left := 2*index + 1
	right := 2*index + 2

	// Find largest value in a[index],a[left] and a[right]
	largest := index
	if left < max && a[left] > a[index] {
		largest = left
	}

	if right < max && a[right] > a[largest] {
		largest = right
	}

	if index != largest {
		a[index], a[largest] = a[largest], a[index]
		heapify(a, largest, max)
	}
}
