package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	// The maximum power of 3 value that integer can hold (3^19)
	return 1162261467%n == 0
}

func main() {
	n := 27
	fmt.Println("Is pwoer of three?", isPowerOfThree(n))
}
