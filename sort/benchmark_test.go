package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	// 1K to 64K
	for s := uint(10); s < uint(17); s = s + 2 {
		for _, bc := range cases {
			b.Run(fmt.Sprintf("%s-%d", bc.Name, 1<<s), func(b *testing.B) {
				b.StopTimer()
				for i := 0; i < b.N; i++ {
					data := make([]int, 1<<s)
					for i := 0; i < len(data); i++ {
						data[i] = rand.Intn(1 << s)
					}
					b.StartTimer()
					bc.Func(data)
					b.StopTimer()
				}
			})
		}

	}
}

func BenchmarkSortSorted(b *testing.B) {
	for _, bc := range cases {
		b.Run(fmt.Sprintf("%s-%d", bc.Name, 1<<s), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := 0; i < len(data); i++ {
					data[i] = i
				}
				b.StartTimer()
				bc.Func(data)
				b.StopTimer()
			}
		})
	}
}
