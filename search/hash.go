package search

type HashTable struct {
	table [][]int
	size  int
}

func NewHashTable(a []int, size int) *HashTable {
	ht := &HashTable{
		table: make([][]int, size),
		size:  size,
	}

	for i := 0; i < len(a); i++ {
		index := ht.hash(a[i])
		ht.table[index%size] = append(ht.table[index%size], a[i])
	}

	return ht
}

func (h *HashTable) hash(x int) int {
	v := x % h.size
	if x < 0 {
		return -v
	}
	return v
}

func (h *HashTable) Search(x int) bool {
	index := h.hash(x)
	l := h.table[index%h.size]
	if len(l) == 0 {
		return false
	}

	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return true
		}
	}

	return false
}
