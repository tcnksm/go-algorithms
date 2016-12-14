package datastructure

type AVLTree struct {
	left, right *AVLTree
	node        interface{}
	less        less

	height int
}

func NewAVLTree(less less) *AVLTree {
	return &AVLTree{
		less: less,
	}
}

// https://gist.github.com/lidashuang/9327631?
func (t *AVLTree) Insert(v interface{}) {
	if t.node == nil {
		t.node = v

		t.right = NewAVLTree(t.less)
		t.left = NewAVLTree(t.less)
		return
	}
}
