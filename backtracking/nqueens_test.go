package backtracking

import (
	"reflect"
	"testing"
)

func TestNQueensRecursive(t *testing.T) {
	nQueenRec4 := [][]int{
		{0, 0, 1, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}

	if a := NQueensRecursive(3); a != nil {
		t.Errorf("NQueens did not work for input 3: %v", a)
	}

	if a := NQueensRecursive(4); !reflect.DeepEqual(a, nQueenRec4) {
		t.Errorf("NQueens did not work for input 4: %v", a)
	}
}

func TestNQueensStack(t *testing.T) {
	nQueensStack4 := [][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 0, 1, 0},
	}

	if a := NQueensStack(3); a != nil {
		t.Errorf("NQueens did not work for input 3: %v", a)
	}

	if a := NQueensStack(4); !reflect.DeepEqual(a, nQueensStack4) {
		t.Errorf("NQueens did not work for input 4: %v", a)
	}
}
