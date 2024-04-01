package main

import "fmt"

func setZeroes(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for row := range rows {
		for i := range matrix[0] {
			matrix[row][i] = 0
		}
	}
	for col := range cols {
		for i := range matrix {
			matrix[i][col] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
