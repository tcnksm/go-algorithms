package datastructure

import "testing"

func TestQueue(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	queue := &Queue{}

	for _, input := range inputs {
		queue.Enqueue(input)
	}

	var v interface{}
	for i := 0; i < len(inputs); i++ {
		var err error
		v, err = queue.Dequeue()
		if err != nil {
			t.Fatal("err:", err)
		}
	}

	if got, want := v.(int), 5; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if _, err := queue.Dequeue(); err == nil {
		t.Fatalf("expect to be failed")
	}
}
