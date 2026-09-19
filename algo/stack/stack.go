package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() T {
	n := len(s.data)
	val := s.data[n-1]
	s.data = s.data[:n-1]
	return val
}

func (s *Stack[T]) Peek() T {
	n := len(s.data)
	return s.data[n-1]
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}
