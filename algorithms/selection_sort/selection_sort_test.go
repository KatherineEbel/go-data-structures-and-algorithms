package selection_sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	s := []int{4, 2, 6, 5, 1, 3}
	result := SelectionSort(s)
	for idx, val := range result {
		if val != idx+1 {
			t.Errorf("unexpected value %d, expected %d", val, idx+1)
		}
	}
}
