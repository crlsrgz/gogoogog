package main

import (
	"fmt"
)

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

func (stack *Stack) push(value int) {

	node := SNode{value, nil}

	stack.length++

	if stack.head == nil && stack.tail == nil {
		stack.head = &node
		// stack.tail = stack.head
		return
	}
	node.prev = stack.head
	stack.head = &node
}
func (stack *Stack) pop() int {

	if stack.length <= 0 {
		stack.length = 0
	} else {
		stack.length -= 1
	}
	if stack.length == 0 {
		head := stack.head
		stack.head = nil
		return head.value
	}
	head := stack.head
	stack.head = head.prev
	return head.value

}

func (stack *Stack) peek() {
	if stack.head != nil {
		fmt.Println(stack.head.value)

	} else {

		fmt.Println("empty")
	}

}
func (stack *Stack) stackLength() {
	fmt.Println("Stack length is: ", stack.length)
}

func main() {

	fmt.Println("///...///")
	myStack := createStack()
	myStack.push(1)
	myStack.push(22)
	myStack.peek()
	myStack.stackLength()

	myStack.push(33)
	myStack.stackLength()
	myStack.peek()
	myStack.pop()
	myStack.peek()
	myStack.stackLength()
	myStack.push(44)
	myStack.push(66)
	myStack.stackLength()
	myStack.peek()
	myStack.pop()
	myStack.peek()
	myStack.stackLength()
	myStack.push(55)
	myStack.peek()
	myStack.stackLength()
	fmt.Println(myStack.head.value, myStack.tail)
}
