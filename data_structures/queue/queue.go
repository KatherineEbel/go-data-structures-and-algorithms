package queue

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](values []T) *Queue[T] {
	return &Queue[T]{data: values}
}

func (s *Queue[T]) Enqueue(element T) *Queue[T] {
	s.data = append(s.data, element)
	return s
}

func (s *Queue[T]) Dequeue() *T {
	if len(s.data) == 0 {
		return nil
	}
	first, rest := s.data[0], s.data[1:]
	s.data = rest
	return &first
}

func (s *Queue[T]) Read() *T {
	if len(s.data) == 0 {
		return nil
	}
	return &s.data[0]
}
