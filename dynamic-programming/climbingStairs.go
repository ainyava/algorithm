package main

import "fmt"

func fib(n int) int {
	prev, next, fib := 0, 1, 1
	for i := 0; i < n; i++ {
		fib += prev
		prev = next
		next = fib
	}
	return prev
}

func climbStairs(n int) int {
	return fib(n + 1)
}

func main() {
	fmt.Println("Ways for climb:", climbStairs(43))
}
