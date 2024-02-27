package main

import (
	"fmt"
)

type List struct {
	head   *Node
	tail   *Node
	length int
}
type NodeData interface {
	string | int64 | uint64 | float64 | bool
}
type Node struct {
	data int
	next *Node
}

func (l *List) createSLL() List {
	return List{}
}

func (l *List) addNode(value int) {
	node := Node{data: value}
	l.length++
	if l.head == nil {
		l.head = &node
		l.tail = &node
		return
	}

	if l.tail != nil {
		l.tail.next = &node
	}

	l.tail = &node

}
func (l *List) removeNode(value int) int {
	if l.head == nil {
		return -1
	}

	if l.head.data == value {
		tmp := l.head
		l.head = l.head.next
		fmt.Println("removed node value is:", tmp.data)

		return tmp.data

	}
	current := l.head
	var tmp int
	for {
		if current == nil {
			fmt.Println("Empty list")
			break
		}
		if current != nil && current.next != nil {
			l.length--
			tmp = current.next.data
			current.next = current.next.next
			break
		}

		current = current.next
	}
	fmt.Println("removed value:", tmp)
	return tmp
}

func (l *List) listLenght() {
	fmt.Printf("The length is %v\n", l.length)
}
func main() {

	myList := List{}
	fmt.Println(myList)
	fmt.Println("///////////////////")
	myList.addNode(2)
	myList.addNode(3)
	myList.addNode(4)
	myList.addNode(5)
	fmt.Println("expe is =", myList.head.next.next.data, "tail is =", myList.tail)
	myList.removeNode(4)
	myList.listLenght()
	fmt.Println("head is =", myList.head.data)
	fmt.Println("tail is =", myList.tail.data)
	fmt.Println("///////////////////")
	myList.removeNode(2)
	myList.listLenght()
	fmt.Println("head is =", myList.head.data)
	fmt.Println("tail is =", myList.tail.data)
	fmt.Println(myList)

}
