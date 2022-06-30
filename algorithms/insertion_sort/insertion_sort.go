package insertion_sort

import (
	"fmt"
)

/*
always starts with the second item
compare it to the item before it
if item before is less than current
switch current with item before

Big O


*/

// InsertionSort sorts a slice using the insertion sort algorithm
func InsertionSort[T int](s []T) {
	// [4 2 6 5 1 3]
	var count int
	for i := 1; i < len(s); i++ {
		count += 1
		for i >= 1 && s[i] < s[i-1] {
			s[i], s[i-1] = s[i-1], s[i]
			count += 1
			i--
		}
	}
	fmt.Println(count)
}

func InsertionSortAlt[T int](s []T) {
	// [4 2 6 5 1 3]
	var count int
	for i := 1; i < len(s); i++ {
		count += 1
		j := i
		for j > 0 {
			count += 1
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
			j -= 1
		}
	}
	fmt.Println(count)
}
