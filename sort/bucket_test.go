package sort

import (
	"sort"
	"testing"
)

func TestBucket(t *testing.T) {
	ints := [...]int{7, 5, 13, 2, 14, 1, 6}
	data := ints
	a := data[:]
	Bucket(a)
	if !sort.IsSorted(sort.IntSlice(a)) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}
