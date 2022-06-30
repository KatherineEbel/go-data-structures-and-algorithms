package bubble_sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := []int{2, 4, 6, 5, 1, 3}
	result := BubbleSort(s)
	for i := 1; i < len(result)+1; i++ {
		if result[i-1] != i {
			t.Errorf("expected %d, got %d", i, result[i])
		}
	}
}
