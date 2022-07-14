package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any](values []T) *Stack[T] {
	return &Stack[T]{data: values}
}

func (s *Stack[T]) Push(element T) *Stack[T] {
	s.data = append(s.data, element)
	return s
}

func (s *Stack[T]) empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() *T {
	if s.empty() {
		return nil
	}
	idx := len(s.data) - 1
	last := s.data[idx]
	s.data = s.data[:idx]
	return &last
}

func (s *Stack[T]) Read() *T {
	if s.empty() {
		return nil
	}
	return &s.data[len(s.data)-1]
}
