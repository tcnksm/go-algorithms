package sort

import "math"

// TODO: Set ideal number of bucket and hash function.
var bucket = math.Pow(2, 10) - 1

func hash(x int) int {
	return x % bucket
}

func Bucket(a []int) {
	b := make(map[int][]int, bucket)
	for i := 0; i < len(a); i++ {
		k := hash(a[i])
		b[k] = append(b[k], a[i])
	}

	extract(b, a)
}

func extract(b map[int][]int, a []int) {

	insertSort := func(a []int) {
		for i := 1; i < len(a); i++ {
			j := i
			for j > 0 && a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				j--
			}
		}
	}

	index := 0
	for i := 0; i < bucket; i++ {
		insertSort(b[i])
		for m := 0; m < len(b[i]); m++ {
			a[index] = b[i][m]
			index++
		}
	}
}
