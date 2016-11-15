package sort

import (
	"sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestBubble(t *testing.T) {
	data := ints
	a := data[:]
	Bubble(a)
	if !sort.IsSorted(sort.IntSlice(a)) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}
