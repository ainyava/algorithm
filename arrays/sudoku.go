package main

import (
	"fmt"
	"slices"
)

func isValidSudoku(board [][]byte) bool {
	var n3 = make(map[int][]byte)
	for row := range board {
		n1 := []byte{}
		n2 := []byte{}
		for col := range board {
			i1 := board[row][col]
			i2 := board[col][row]

			rowp := row / 3
			colp := col / 3
			key := 7*rowp + 5*colp

			if slices.Contains(n1, i1) || slices.Contains(n2, i2) || slices.Contains(n3[key], i1) {
				return false
			}
			if i1 != '.' {
				n1 = append(n1, i1)
				n3[key] = append(n3[key], i1)
			}
			if i2 != '.' {
				n2 = append(n2, i2)
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println("Valid board?", isValidSudoku(board))
}
