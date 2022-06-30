package insertion_sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	s := []int{4, 2, 6, 5, 1, 3}
	InsertionSort(s)
	for i, val := range s {
		if val != i+1 {
			t.Errorf("unexpected value, got %d, wanted %d", val, i+1)
		}
	}
}

func TestInsertionSortSortedArray(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	InsertionSort(s)
	for i, val := range s {
		if val != i+1 {
			t.Errorf("unexpected value, got %d, wanted %d", val, i+1)
		}
	}
}

func TestInsertionSort2(t *testing.T) {
	s := []int{2, 5, 6, 1, 3, 4}
	InsertionSort(s)
	for i, val := range s {
		if val != i+1 {
			t.Errorf("unexpected value, got %d, wanted %d", val, i+1)
		}
	}
}

func TestInsertionSortAlt(t *testing.T) {
	s := []int{2, 5, 6, 1, 3, 4}
	InsertionSortAlt(s)
	for i, val := range s {
		if val != i+1 {
			t.Errorf("unexpected value, got %d, wanted %d", val, i+1)
		}
	}
}
