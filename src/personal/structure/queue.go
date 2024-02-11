package main

import "fmt"

type QueueNode struct {
	value int
	prev  *QueueNode
}
type Stack struct {
	head   *QueueNode
	tail   int
	length int
}

func createStack(value int) *Stack {
	node := new(QueueNode)
	node.value = value

	stack := new(Stack)

	stack.length++
	if stack.head == nil {
		stack.head = node
		return stack
	}
	node.prev = stack.head
	stack.head = node
	return stack
}
func main() {
	fmt.Println(createStack(5))
	fmt.Println(createStack(8))
}
