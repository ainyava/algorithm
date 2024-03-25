package main

import "fmt"

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {
	triangle := generate(5)
	for _, row := range triangle {
		fmt.Println(row)
	}
}
