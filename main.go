package main

import (
	"fmt"
	Container "learning-go/Containers/Stack"
)

func main() {

	stack := &Container.Stack{Top: nil}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
