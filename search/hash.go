package search

import (
	"fmt"
	"hash/fnv"
	"math"
)

var tableLen int

var (
	defaultTableLen = int(math.Pow(2, 10)) - 1
)

func init() {
	tableLen = defaultTableLen
}

// HashTable construct hash table with size n.
func HashTable(a []int) [][]int {
	table := make([][]int, tableLen)
	h := fnv.New32()
	for _, v := range a {
		h.Write([]byte(fmt.Sprintf("%d", v)))
		index := int(h.Sum32())
		h.Reset()
		table[index%tableLen] = append(table[index%tableLen], v)
	}

	return table
}

// Hash search searches target x is in HashTable or not
func Hash(table [][]int, x int) bool {
	h := fnv.New32()
	h.Write([]byte(fmt.Sprintf("%d", x)))
	index := int(h.Sum32())

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
