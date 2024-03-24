package main

import "fmt"

func hammingDistance(x int, y int) int {
	count := 0
	for x > 0 || y > 0 {
		count += (x & 1) ^ (y & 1)
		x >>= 1
		y >>= 1
	}
	return count
}

func main() {
	x := 3
	y := 1
	fmt.Println("Hamming weight:", hammingDistance(x, y))
}
