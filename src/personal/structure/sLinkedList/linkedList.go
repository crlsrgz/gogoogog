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
	string | int | int64 | uint64 | float64 | bool
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
func (l *List) removeByValue(value int) (int, bool) {
	if l.head == nil {
		return -1, false
	}
	//: Is the head

	if l.head.data == value {
		l.length--
		tmp := l.head
		l.head = l.head.next
		fmt.Println("removed node value is:", tmp.data)

		return tmp.data, true
	}
	//: Is after the head?
	current := l.head
	var tmp int

	for {
		if current == nil {
			fmt.Println("Empty list")
			return -1, false
		}
		//: Is the Tail?

		if current.next.data == value && l.tail.data == value {
			l.length--
			tmp = l.tail.data
			current.next = nil
			l.tail = current
			return tmp, true
		}

		if current.next.data == value {
			l.length--
			tmp = current.next.data
			// fmt.Println("current", current.data)
			// fmt.Println("current.next.next", current.next.next.data)
			current.next = current.next.next

			return tmp, true
		}
		if current.next == l.tail {
			fmt.Println("nomore")
			return -1, false
		}
		current = current.next
	}
	// fmt.Println("removed value:", tmp)

}

func (l *List) listLenght() int {
	return l.length
}

func (l *List) printValues() []int {
	if l.length <= 0 {
		fmt.Println("The list is empty")
	}
	values := []int{}
	current := l.head

	for {
		if current == nil {
			// fmt.Println("List values:", values)
			return values
		}
		values = append(values, current.data)
		current = current.next
	}
}

// /TEMP

func main() {

	myList := List{}
	fmt.Println(myList)
	myList.printValues()
	fmt.Println("///////////////////")
	myList.addNode(2)
	myList.addNode(3)
	myList.addNode(4)
	myList.addNode(5)
	myList.addNode(6)
	myList.printValues()
	fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	fmt.Println("///////////////////")

	if val, ok := myList.removeByValue(3); ok {

		fmt.Println("is here ->", val)
	} else {
		fmt.Println("not here ->", ok)
	}
	fmt.Println("///////////////////")
	fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	fmt.Println("///////////////////")
	// myList.removeByValue(5)
	// myList.printValues()
	// fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	// fmt.Println("///////////////////")
	// myList.removeByValue(6)
	// myList.printValues()
	// fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	// myList.removeByValue(2)
	// myList.printValues()
	// fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	// fmt.Println("///////////////////")
	// myList.addNode(2)
	// myList.addNode(3)
	// myList.addNode(5)
	// myList.addNode(5)
	// myList.addNode(6)
	// myList.removeByValue(4)
	// myList.printValues()
	// fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())
	// fmt.Println("///////////////////")
	// myList.removeByValue(4)
	// myList.removeByValue(8)
	// myList.printValues()
	// fmt.Printf("head: %v, tail: %v, lenght: %v\n", myList.head.data, myList.tail.data, myList.listLenght())

}
