package stack_test

import (
	"testing"

	"data-structures-and-algorithms/data_structures/stack"
)

func TestStackRead(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	s := stack.NewStack(vals)
	a := s.Read()
	if *a != 5 {
		t.Errorf("expected %d, got %d", 5, a)
	}
}

func TestStackPush(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	s := stack.NewStack(vals)
	s.Push(6)
	a := s.Read()
	if *a != 6 {
		t.Errorf("expected %d, got %d", 6, a)
	}
}

func TestStackPop(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	s := stack.NewStack(vals)
	a := s.Pop()
	if *a != 5 {
		t.Errorf("expected %d, got %d", 5, a)
	}
	a = s.Read()
	if *a != 4 {
		t.Errorf("expected %d, got %d", 4, a)
	}
}
