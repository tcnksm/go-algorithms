package datastructure

import (
	"sort"
	"testing"
)

func TestBinary(t *testing.T) {
	less := func(x, y interface{}) bool {
		i, j := x.(int), y.(int)
		return i < j
	}

	tree := NewBinaryTree(less)
	for _, i := range []int{1, 10, 14, 3, 56} {
		tree.Insert(i)
	}

	var res []int
	tree.Walk(func(v interface{}) {
		i := v.(int)
		res = append(res, i)
	})

	if !sort.IsSorted(sort.IntSlice(res)) {
		t.Fatalf("expect to be sorted: %v", res)
	}
}
