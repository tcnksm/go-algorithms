package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	cases := []struct {
		Name string
		Func func([]int)
	}{
		{
			"Standard",
			sort.Ints,
		},

		{
			"Insert",
			Insert,
		},

		{
			"Bubble",
			Bubble,
		},
	}
	sizes := []int{1000, 5000, 10000}
	for _, size := range sizes {
		for _, bc := range cases {
			b.Run(fmt.Sprintf("%s-%d", bc.Name, size), func(b *testing.B) {
				b.StopTimer()
				for i := 0; i < b.N; i++ {
					data := make([]int, size)
					for i := 0; i < len(data); i++ {
						data[i] = rand.Intn(size)
					}
					b.StartTimer()
					bc.Func(data)
					b.StopTimer()
				}
			})
		}

	}
}
