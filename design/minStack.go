package main

import (
	"fmt"
)

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if this.min[len(this.min)-1] == this.stack[len(this.stack)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	obj.Pop()
	fmt.Println("Stack:", obj)
}
