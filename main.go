package main

import (
	"algorithms/backtracking"
)

func main() {

	// quick sort implementation
	// fmt.Println(customsort.QuickSort([]int{50, 23, 9, 18, 61, 32}))
	// fmt.Println(customsort.QuickSort([]int{1, 2, 3, 4, 5, 5}))

	// // quick select implementation
	// fmt.Println(customselect.QuickSelect([]int{50, 23, 9, 18, 61, 32}, 3))
	// fmt.Println(customselect.QuickSelect([]int{1, 2, 3, 4, 5, 5}, 6))

	// nqueens implementation
	// backtracking.NQueens(8)

	// hamiltonian cycle
	graph := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	// backtracking.Hamiltoniancycle(graph)
	backtracking.DFS(graph)

}
