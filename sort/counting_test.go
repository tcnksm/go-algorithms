package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestCounting(t *testing.T) {
	k := 10
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = rand.Intn(k)
	}

	Counting(a, k)
	if !sort.IsSorted(sort.IntSlice(a)) {
		t.Errorf("got %v", a)
	}
}
