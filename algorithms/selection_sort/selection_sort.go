package selection_sort

/*
Have to keep track of indexes
var min set to index first item
iterate over array if current value less than current min store that index to min

once at end of array -
take item at min index and swap with start index if minIdx is not equal to current
item at start index is now sorted

BigO


*/

// SelectionSort sorts an integer array using selection sort algorithm
func SelectionSort[T int](s []T) []T {
	// [4 2 6 5 1 3]
	for idx, _ := range s {
		minIdx := idx
		for i := idx + 1; i < len(s); i++ {
			if s[i] < s[minIdx] {
				minIdx = i
			}
		}
		if minIdx != idx {
			s[idx], s[minIdx] = s[minIdx], s[idx]
		}
	}
	return s
}
