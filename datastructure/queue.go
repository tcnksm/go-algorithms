package datastructure

import "fmt"

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, fmt.Errorf("empty")
	}

	v := q.data[0]
	q.data = q.data[1:]
	return v, nil
}
