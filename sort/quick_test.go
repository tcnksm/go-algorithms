package sort

import (
	"sort"
	"testing"
)

func TestQuickSimple(t *testing.T) {
	data := ints
	a := data[:]
	a = quickSimple(a)
	if !sort.IsSorted(sort.IntSlice(a)) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}
