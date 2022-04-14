package main

import (
	"algorithms/backtracking"
)

func main() {

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
