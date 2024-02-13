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

func (stack *Stack) addNode(value int) {

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
	myStack.addNode(1)
	myStack.addNode(22)
	myStack.peek()
	myStack.stackLength()

	myStack.addNode(33)
	myStack.peek()
	myStack.stackLength()
	fmt.Println(myStack.head, myStack.tail)
}
