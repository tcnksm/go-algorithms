package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSort1K(b *testing.B) {
	for _, bc := range cases {
		b.Run(fmt.Sprintf("%s", bc.Name), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := 0; i < len(data); i++ {
					data[i] = rand.Intn(1 << 10)
				}
				b.StartTimer()
				bc.Func(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSortSorted(b *testing.B) {
	for _, bc := range cases {
		b.Run(fmt.Sprintf("%s", bc.Name), func(b *testing.B) {
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
