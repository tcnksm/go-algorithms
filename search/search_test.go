package search

import (
	"sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestSequential(t *testing.T) {
	data := ints
	a := data[:]
	if got := Sequential(a, 42); !got {
		t.Errorf("Search=%v, want=true", got)
	}

	if got := Sequential(a, 43); got {
		t.Errorf("Search=%v, want=false", got)
	}
}

func TestBinary(t *testing.T) {
	data := ints
	a := data[:]
	sort.Ints(a)
	if got := Binary(a, 42); !got {
		t.Errorf("Search=%v, want=true", got)
	}

	if got := Binary(a, 43); got {
		t.Errorf("Search=%v, want=false", got)
	}
}
