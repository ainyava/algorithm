package main

import "fmt"

func isBadVersion(version int) bool {
	return version >= 4
}

func firstBadVersion(n int) int {
	low, high := 0, n
	for low < high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
func main() {
	n := 5
	fmt.Println("First bad version is:", firstBadVersion(n))
}
