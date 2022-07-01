package merge_sort

import (
	"testing"
)

func TestMerge1(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s1 := []int{1}
	s2 := []int{2}
	result := Merge(s1, s2)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}

func TestMerge(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s1 := []int{1, 3, 7, 8}
	s2 := []int{2, 4, 5, 6}
	result := Merge(s1, s2)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}

func TestMerge2(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s1 := []int{1, 3, 7}
	s2 := []int{2, 4, 5, 6}
	result := Merge(s1, s2)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}

func TestMerge3(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s1 := []int{1, 3, 6, 7}
	s2 := []int{2, 4, 5}
	result := Merge(s1, s2)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}

func TestMerge4(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s1 := []int{1, 3, 6, 9}
	s2 := []int{2, 4, 5, 7, 8}
	result := Merge(s1, s2)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}

func TestMergeSort(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s := []int{2, 6, 1, 4, 5, 7, 3, 8}
	result := MergeSort(s)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}

func TestMergeSort2(t *testing.T) {
	// [1 3 7 8] [2 4 5 6]
	s := []int{2, 6, 1, 4, 5, 7, 3}
	result := MergeSort(s)
	for idx, val := range result {
		if idx+1 != val {
			t.Errorf("unexpected value %d, wanted %d", val, idx+1)
		}
	}
}
