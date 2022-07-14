package quick_sort

import (
	"fmt"
	"testing"
)

func TestPivot(t *testing.T) {
	s := []int{4, 6, 1, 7, 3, 2, 5}
	result := Pivot(s)
	if result != 3 {
		t.Errorf("expected %d, got %d", 3, result)
	}
}

func TestPivot2(t *testing.T) {
	s := []int{4, 8, 2, 1, 5, 7, 6, 3}
	result := Pivot(s)
	if result != 3 {
		t.Errorf("expected %d, got %d", 3, result)
	}
}

func TestPivot3(t *testing.T) {
	s := []int{3, 8, 2, 1, 5, 7, 6, 4}
	result := Pivot(s)
	if result != 2 {
		t.Errorf("expected %d, got %d", 2, result)
	}
}

func TestPivot4(t *testing.T) {
	s := []int{5, 8, 2, 1, 3, 7, 6, 4}
	result := Pivot(s)
	if result != 4 {
		t.Errorf("expected %d, got %d", 4, result)
	}
}

func TestQuickSort(t *testing.T) {
	s := []int{4, 6, 1, 7, 3, 2, 5}
	QuickSort(s)
	fmt.Print(s)
	for i, val := range s {
		if val != i+1 {
			t.Errorf("expected %d, got %d", i, val)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	s := []int{6, 2, 5, 4, 1, 7, 3}
	QuickSort(s)
	fmt.Print(s)
	for i, val := range s {
		if val != i+1 {
			t.Errorf("expected %d, got %d", i, val)
		}
	}
}
