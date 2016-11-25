package search

import (
	"fmt"
	"math/rand"
	"testing"
)

// randomArray returns random array but inject the given
// value x at index p.
func randomArray(n, p, x int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		if i == p {
			a[i] = x
			continue
		}
		a[i] = rand.Intn(n)
	}

	return a
}

func BenchmarkSequential(b *testing.B) {

	for k := uint(8); k <= 16; k++ {
		l := 1 << k
		b.Run(fmt.Sprintf("%d", l), func(b *testing.B) {
			b.StopTimer()
			x := 124
			a := randomArray(l, l/2, x)
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				if !Sequential(a, x) {
					b.Fatal("expect to be found")
				}
			}
		})
	}
}
