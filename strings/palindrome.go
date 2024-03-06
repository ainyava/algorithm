package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	var onlyLetters []rune
	for _, ch := range s {
		if ch >= 65 && ch <= 90 {
			ch |= 32
		}
		if (ch >= 97 && ch <= 122) || (ch >= 48 && ch <= 57) {
			onlyLetters = append(onlyLetters, ch)
		}
	}
	for i := 0; i < len(onlyLetters)/2; i++ {
		if onlyLetters[i] != onlyLetters[len(onlyLetters)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println("Palindrome?", isPalindrome(s))
}
