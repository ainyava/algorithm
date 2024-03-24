package main

import "fmt"

func reverseBits(n uint32) uint32 {
	var reversed uint32 = 0
	for range 32 {
		reversed <<= 1
		if n&1 == 1 {
			reversed ^= 1
		}
		n >>= 1
	}
	return reversed
}

func main() {
	var n uint32 = 43261596
	fmt.Println("Reversed bits:", reverseBits(n))
}
