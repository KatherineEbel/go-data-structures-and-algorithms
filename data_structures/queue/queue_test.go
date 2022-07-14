package queue_test

import (
	"testing"

	"data-structures-and-algorithms/data_structures/queue"
)

func TestNewQueue(t *testing.T) {
	q := queue.NewQueue([]int{})
	q.Enqueue(1)
	r := q.Read()
	if *r != 1 {
		t.Errorf("expected %d, got %d", 1, *r)
	}
	r = q.Dequeue()
	if *r != 1 {
		t.Errorf("expected %d, got %d", 1, *r)
	}

	r = q.Read()
	if r != nil {
		t.Errorf("expected %d, got %d", 1, *r)
	}

	r = q.Dequeue()
	if r != nil {
		t.Errorf("expected %d, got %d", 1, *r)
	}
}
