package bloom

import "testing"

func TestBloomFilter(t *testing.T) {

	inputs := []string{"x", "y", "z"}

	BF := New(100, 8000)
	for _, input := range inputs {
		BF.Add([]byte(input))
	}

	if got := BF.Contains([]byte("a")); got {
		t.Fatalf("Contains('a') = %v, want false", got)
	}

	if got := BF.Contains([]byte("x")); !got {
		t.Fatalf("Contains('x') = %v, want true", got)
	}
}
