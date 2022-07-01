package merge_sort

/*

Merge Sort Time Complexity is O(n log n)
Space complexity is O(n)
*/

// Merge takes two sorted arrays and combines them
func Merge(s1, s2 []int) []int {
	// [1 3 7 8] [2 4 5 6]
	var merged []int
	var i, j int
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			merged = append(merged, s1[i])
			i++
		} else {
			merged = append(merged, s2[j])
			j++
		}
	}
	merged = append(merged, s1[i:]...)
	merged = append(merged, s2[j:]...)
	return merged
}

func MergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	left := MergeSort(s[:len(s)/2])
	right := MergeSort(s[len(s)/2:])
	return Merge(left, right)
}
