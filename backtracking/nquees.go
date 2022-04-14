package backtracking

import "fmt"

var chessBoard [][]int

func NQueens(n int) {

	chessBoard = make([][]int, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]int, n)
		for j := 0; j < n; j++ {
			chessBoard[i][j] = 0
		}
	}
	// printBoard()

	if nqueensstack(n) {
		printBoard()
	} else {
		fmt.Println("No solution available")
	}
}

// stack implementation of nqueens
func nqueensstack(n int) bool {
	count := 0
	stack := make([][2]int, 0)

	// add the first level children to the stack
	for i := 0; i < n; i++ {
		stack = append(stack, [2]int{i, 0})
	}

	//pop stack
	// if the popped element is -1, -1 (child marker, can also do it with the children) it means that the current parent did not lead to a solution and
	// hence remove parent, decrease count and unmark the chess board
	// if the popped element is valid
	// add it back to the stack
	// add a child marker and add all its children
	// increment count and mark the chessboard
	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		row, col := popped[0], popped[1]
		if row == -1 && col == -1 {
			popped = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			row, col := popped[0], popped[1]
			chessBoard[row][col] = 0
			count--
		} else if valid(row, col, n) {
			chessBoard[row][col] = 1
			count++
			if count == n {
				return true
			}
			stack = append(stack, [2]int{row, col})
			stack = append(stack, [2]int{-1, -1})
			for i := 0; i < n; i++ {
				stack = append(stack, [2]int{i, col + 1})
			}
		}
	}

	return false
}

// recursive implementation of nqueens
func nqueens(col int, n int) bool {
	if col >= n {
		return true
	}

	for row := 0; row < n; row++ {
		if valid(row, col, n) {
			// place a queen
			chessBoard[row][col] = 1

			if nqueens(col+1, n) {
				return true
			}

			chessBoard[row][col] = 0
		}
	}

	return false
}

func valid(row, col, n int) bool {
	for i := 0; i < col; i++ {
		if chessBoard[row][i] == 1 {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if chessBoard[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	for i, j := row+1, col-1; i < n && j >= 0; {
		if chessBoard[i][j] == 1 {
			return false
		}
		i++
		j--
	}

	return true
}

func printBoard() {
	for _, rows := range chessBoard {
		for _, v := range rows {
			if v == 1 {
				fmt.Printf("*")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
