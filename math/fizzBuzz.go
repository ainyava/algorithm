package main

import "fmt"

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		item := ""
		if i%3 == 0 {
			item += "Fizz"
		}
		if i%5 == 0 {
			item += "Buzz"
		} else if i%3 != 0 {
			item += fmt.Sprintf("%v", i)
		}
		ans[i-1] = item
	}
	return ans
}

func main() {
	n := 3
	fmt.Println("Result:", fizzBuzz(n))
}
