package sort

import (
	"sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

var cases = []struct {
	Name     string
	Func     func([]int)
	SkipTest bool
}{

	{
		"Quick",
		Quick,
		false,
	},

	{
		"Shell",
		Shell,
		false,
	},

	{
		"Heap",
		Heap,
		false,
	},

	{
		"Bucket",
		Bucket,
		true,
	},

	{
		"Std",
		sort.Ints,
		true,
	},

	{
		"Insert",
		Insert,
		false,
	},

	{
		"Select",
		Selection,
		false,
	},

	{
		"Bubble",
		Bubble,
		false,
	},
}

func TestSort(t *testing.T) {
	for _, tc := range cases {
		if tc.SkipTest {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			data := ints
			a := data[:]
			tc.Func(a)
			if !sort.IsSorted(sort.IntSlice(a)) {
				t.Errorf("sorted %v", ints)
				t.Errorf("   got %v", data)
			}
		})
	}
}
