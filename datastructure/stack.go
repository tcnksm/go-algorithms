package datastructure

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, fmt.Errorf("empty")
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return v, nil
}
