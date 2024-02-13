package main

import "fmt"

type SNode struct {
	value int
	prev  *SNode
}

type Stack struct {
	head   *SNode
	tail   *SNode
	length int
}

func createStack() Stack {
	var stack = new(Stack)
	return *stack
}

func (stack *Stack) addNode(value int) *Stack {

	node := SNode{value, nil}

	stack.length++

	if stack.head == nil && stack.tail == nil {
		stack.head = &node
		stack.tail = stack.head
	}
	return stack
}

func main() {

	fmt.Println("///...///")
	myStack := createStack()
	myStack.addNode(3)
	fmt.Println(myStack.head, myStack.tail)
}
