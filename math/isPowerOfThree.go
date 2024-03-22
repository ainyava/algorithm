package main

import "fmt"

func isPowerOfThree(n int) bool {
	for bases := 1; bases <= n; bases *= 3 {
		if bases == n {
			return true
		}
	}
	return false
}

func main() {
	n := 27
	fmt.Println("Is pwoer of three?", isPowerOfThree(n))
}
