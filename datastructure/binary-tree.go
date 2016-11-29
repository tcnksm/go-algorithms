package datastructure

type less func(x, y interface{}) bool

type BinaryTree struct {
	left, right *BinaryTree
	node        interface{}

	less less
}

func NewBinaryTree(less less) *BinaryTree {
	return &BinaryTree{
		less: less,
	}
}

// Insert inserts element v into BinaryTree
func (t *BinaryTree) Insert(v interface{}) {
	if t.node == nil {
		t.node = v
		t.left = NewBinaryTree(t.less)
		t.right = NewBinaryTree(t.less)
		return
	}

	if t.less(v, t.node) {
		t.left.Insert(v)
		return
	}

	t.right.Insert(v)
}

func (t *BinaryTree) Search(v interface{}) bool {
	if t.node == nil {
		return false
	}

	if t.node == v {
		return true
	}

	if t.less(v, t.node) {
		return t.left.Search(v)
	}

	return t.right.Search(v)
}
