package search

import "github.com/tcnksm/go-algorithms/datastructure"

// Binary Search Tree is very suited when you need to
// add (insert) new elements very often.
func NewBinaryTree(a []int) *datastructure.BinaryTree {
	less := func(x, y interface{}) bool {
		i, j := x.(int), y.(int)
		return i < j
	}

	tree := datastructure.NewBinaryTree(less)
	for i := 0; i < len(a); i++ {
		tree.Insert(a[i])
	}

	return tree
}
