package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	min := math.MaxInt32
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		profit := prices[i] - min
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Total profit: ", maxProfit(prices))
}
