package search

import "github.com/tcnksm/go-algorithms/datastructure/tree"

// Binary Search Tree is very suited when you need to
// add (insert) new elements very often.
func NewBinaryTree(a []int) *tree.Binary {
	less := func(x, y interface{}) bool {
		i, j := x.(int), y.(int)
		return i < j
	}

	tree := tree.NewBinary(less)
	for i := 0; i < len(a); i++ {
		tree.Insert(a[i])
	}

	return tree
}
