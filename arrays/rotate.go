package main

import (
	"fmt"
)

func rotateMatrix(matrix [][]int) {
	clone := make([][]int, len(matrix))
	for i := range matrix {
		clone[i] = make([]int, len(matrix[i]))
		copy(clone[i], matrix[i])
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][len(matrix)-i-1] = clone[i][j]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateMatrix(matrix)
	fmt.Println("Rotated matrix?", matrix)
}
