package search

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
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

func BenchmarkBinary(b *testing.B) {
	l := 1 << 20
	b.Run(fmt.Sprintf("%d", l), func(b *testing.B) {
		b.StopTimer()
		x := rand.Intn(l)
		a := randomArray(l, rand.Intn(l-1), x)
		sort.Ints(a)
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			if !Binary(a, x) {
				b.Fatalf("expect %d to be found", x)
			}
			b.StopTimer()
		}
	})
}

func BenchmarkHash(b *testing.B) {
	for k := uint(8); k <= 16; k++ {
		l := 1 << k
		b.Run(fmt.Sprintf("%d", l), func(b *testing.B) {
			b.StopTimer()
			x := rand.Intn(l)
			a := randomArray(l, l/2, x)
			b.StartTimer()

			ht := NewHashTable(a, int(math.Pow(2, 10))-1)
			for i := 0; i < b.N; i++ {
				if ht.Search(x) {
					b.Fatal("expect to be found")
				}
			}
		})
	}
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
