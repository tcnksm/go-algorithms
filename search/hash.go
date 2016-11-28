package search

import "math"

var tableLen int

var (
	defaultTableLen = int(math.Pow(2, 10)) - 1
)

func init() {
	tableLen = defaultTableLen
}

func hash(x int) int {
	v := x % tableLen
	if x < 0 {
		return -v
	}
	return v
}

// HashTable construct hash table with size n.
func HashTable(a []int) [][]int {
	table := make([][]int, tableLen)
	for _, v := range a {
		index := hash(v)
		table[index%tableLen] = append(table[index%tableLen], v)
	}

	return table
}

// Hash search searches target x is in HashTable or not
func Hash(table [][]int, x int) bool {
	index := hash(x)
	l := table[index%tableLen]
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
