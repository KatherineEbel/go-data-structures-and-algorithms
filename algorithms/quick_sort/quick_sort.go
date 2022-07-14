package quick_sort

/*
pivot points to first
swap set to pivot
i points to second
  start by setting second el to i
loop through until i < pivot
  - move swap forward and swap i and swap
  - at end of loop swap pivot and swap
return swap
*/

// Pivot is a helper function for quick sort
func Pivot(s []int) int {
	// [4 6 1 7 3 2 5]
	start := 0
	pivot := s[start]
	pivotIdx := 0
	for i := start + 1; i < len(s); i++ {
		if pivot > s[i] {
			pivotIdx++
			s[i], s[pivotIdx] = s[pivotIdx], s[i]
		}
	}
	s[start], s[pivotIdx] = s[pivotIdx], s[start]
	return pivotIdx
}

func QuickSort(s []int) {
	if len(s) <= 1 {
		return
	}
	pivotIndex := Pivot(s)
	QuickSort(s[0:pivotIndex])
	QuickSort(s[pivotIndex+1:])
	return
}
