package main

import (
	"fmt"
	"slices"
)

type Stack []rune

func (stack *Stack) push(item rune) {
	*stack = append(*stack, item)
}

func (stack *Stack) pop() rune {
	if len(*stack) == 0 {
		return ' '
	}
	index := len(*stack) - 1
	el := (*stack)[index]
	*stack = (*stack)[:index]
	return el
}

func isValid(s string) bool {
	var stack Stack
	for _, i := range s {
		if slices.Contains([]rune{'(', '{', '['}, i) {
			stack.push(i)
		}
		if i == '}' && stack.pop() != '{' {
			return false
		}
		if i == ')' && stack.pop() != '(' {
			return false
		}
		if i == ']' && stack.pop() != '[' {
			return false
		}
	}
	if stack.pop() == ' ' {
		return true
	}
	return false
}

func main() {
	s := "(){}[]()"
	fmt.Println("Is Valid?", isValid(s))
}
