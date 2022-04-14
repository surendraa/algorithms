package customsort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// unsorted input
	a := QuickSort([]int{50, 23, 9, 18, 61, 32})
	if !sort.IntsAreSorted(a) {
		t.Errorf("Expected sorted array but got %v", a)
	}

	// sorted repeated input
	a = QuickSort([]int{1, 2, 3, 4, 5, 5})
	if !sort.IntsAreSorted(a) {
		t.Errorf("Expected sorted array but got %v", a)
	}
}
