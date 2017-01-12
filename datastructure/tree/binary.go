package tree

type less func(x, y interface{}) bool

type Binary struct {
	left, right *Binary
	node        interface{}

	less less
}

func NewBinary(less less) *Binary {
	return &Binary{
		less: less,
	}
}

// Insert inserts element v into BinaryTree
func (t *Binary) Insert(v interface{}) {
	if t.node == nil {
		t.node = v
		t.left = NewBinary(t.less)
		t.right = NewBinary(t.less)
		return
	}

	if t.less(v, t.node) {
		t.left.Insert(v)
		return
	}

	t.right.Insert(v)
}

func (t *Binary) Search(v interface{}) bool {
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

func (t *Binary) Walk(walkFn func(v interface{})) {
	if t.node == nil {
		return
	}

	t.left.Walk(walkFn)
	walkFn(t.node)
	t.right.Walk(walkFn)
}
