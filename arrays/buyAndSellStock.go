package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	prices := []int{2, 1, 2, 1, 0, 1, 2}
	fmt.Println("Total profit: ", maxProfit(prices))
}
