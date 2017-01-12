package datastructure

import "testing"

func TestStack(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	stack := &Stack{}

	for _, input := range inputs {
		stack.Push(input)
	}

	var v interface{}
	for i := 0; i < len(inputs); i++ {
		var err error
		v, err = stack.Pop()
		if err != nil {
			t.Fatal("err:", err)
		}
	}

	if got, want := v.(int), 1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if _, err := stack.Pop(); err == nil {
		t.Fatalf("expect to be failed")
	}

}
