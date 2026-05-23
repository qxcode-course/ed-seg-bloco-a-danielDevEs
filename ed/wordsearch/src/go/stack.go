package main

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	idx := len(s.data)-1
	value := s.data[idx]

	var zero T
	s.data[idx] = zero

	s.data = s.data[:idx]
	return value
}
