package customselect

import (
	"testing"
)

func TestCustomSelect(t *testing.T) {

	// unsorted input
	a := QuickSelect([]int{50, 23, 9, 18, 61, 32}, 3)
	if a != 23 {
		t.Errorf("Expected %v but got %v", 23, a)
	}

	// sorted repeated input
	a = QuickSelect([]int{1, 2, 3, 4, 5, 5}, 6)
	if a != 5 {
		t.Errorf("Expected %v but got %v", 5, a)
	}
}
