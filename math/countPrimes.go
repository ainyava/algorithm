package main

import "fmt"

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	is_not_prime := make([]bool, n)
	is_not_prime[0], is_not_prime[1] = true, true
	// Sieve of Eratosthenes
	i := 2
	for i*i <= n {
		if is_not_prime[i] {
			i++
			continue
		}
		j := 2 * i
		for j < n {
			is_not_prime[j] = true
			j += i
		}
		i++
	}
	// Count
	count := 0
	for _, item := range is_not_prime {
		if !item {
			count++
		}
	}
	return count
}

func main() {
	n := 10
	fmt.Println("Count of primes:", countPrimes(n))
}
