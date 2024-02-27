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
		l.length--
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
		if l.tail.data == value {
			l.length--
			fmt.Println("current", current)
			tmp = l.tail.data
			l.tail = current.next
			current.next = nil
			break
		}
		if current.next.data == value {
			l.length--
			tmp = current.next.data
			fmt.Println("current", current)
			fmt.Println("current.next.next", current.next.next)
			current = current.next.next

			break
		}

		current = current.next
	}
	fmt.Println("removed value:", tmp)
	return tmp
}

func (l *List) listLenght() int {
	// fmt.Printf("The length is %v\n", l.length)
	return l.length
}
func main() {

	myList := List{}
	fmt.Println(myList)
	fmt.Println("///////////////////")
	myList.addNode(2)
	myList.addNode(3)
	myList.addNode(4)
	myList.addNode(5)
	myList.addNode(6)
	fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	fmt.Println(myList.head.next.next)
	myList.removeNode(3)
	fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	fmt.Println("///////////////////")
	myList.removeNode(5)
	fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	fmt.Println("///////////////////")
	myList.removeNode(6)
	fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())

}
