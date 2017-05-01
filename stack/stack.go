package stack

import (
	"errors"
)

type node struct {
	number int
	down   *node
}

type Stack struct {
	root *node
}

func (s *Stack) Peek() (int, error) {
	if s.root == nil {
		return 0, errors.New("Nothing to peek")
	}

	return s.root.number, nil
}

func (s *Stack) Push(number int) {

	tmp := &node{number: number, down: s.root}
	s.root = tmp
}

func (s *Stack) Pop() (int, error) {
	if s.root == nil {
		return 0, errors.New("Nothing to pop")
	}

	ret := s.root
	s.root = ret.down

	return ret.number, nil
}
