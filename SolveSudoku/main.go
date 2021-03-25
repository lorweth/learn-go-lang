package main

import "fmt"

// checkValid func: to check param k has duplicate in horizontal/vertical/set 9
func checkValid(S [9][9]int, x, y, k int) bool {
	// check horizontal
	for i := 0; i < 9; i++ {
		if S[x][i] == k {
			return false
		}
	}
	// check vertical
	for i := 0; i < 9; i++ {
		if S[i][y] == k {
			return false
		}
	}
	// check set 9
	a, b := x/3, y/3

	for i := a * 3; i < a*3+3; i++ {
		for j := b * 3; j < b*3+3; j++ {
			if S[i][j] == k {
				return false
			}
		}
	}
	return true
}

func printSolution(S [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%v\t", S[i][j])
		}
		fmt.Println()
	}
}

func solveSudoku(S [9][9]int, x, y int) {
	if y == 9 {
		if x == 8 {
			// print solution
			printSolution(S)
		} else {
			solveSudoku(S, x+1, 0)
		}
	} else if S[x][y] == 0 {
		for k := 0; k <= 9; k++ {
			if checkValid(S, x, y, k) {
				S[x][y] = k
				solveSudoku(S, x, y+1)
				S[x][y] = 0
			}
		}
	} else {
		solveSudoku(S, x, y+1)
	}
}

func main() {
	sudoku1 := [9][9]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	// start at 0,0
	solveSudoku(sudoku1, 0, 0)

}
