package main

import "fmt"
type QNodeValuer interface {}

type QNode struct {
	value QNodeValuer
	next *QNode
}

type Queue struct{
	head *QNode
	tail *QNode
	length int
}
func (q *Queue) enqueue(value QNodeValuer) {
	node := QNode{}
	node.value = value
	q.length++
	

	if q.tail == nil {
		q.head =  &node
		q.tail = q.head 
		return
	}

	q.tail.next = &node
	q.tail = &node

}



func main() {
	fmt.Println("Init")
	
	intNode := QNode{10, nil}
	stringNode := QNode{"ten", nil}
	
	fmt.Println("testing nodes", intNode, stringNode)

	fmt.Println("////");
	myQueue := Queue{}
	fmt.Println(myQueue)
	fmt.Println("Queue length: ", myQueue.length)
	fmt.Println("////");
	myQueue.enqueue(4)
	fmt.Println("Queue length: ", myQueue.length)
	fmt.Println(myQueue.tail)
	fmt.Println("////+1");
	myQueue.enqueue("four")
	fmt.Println("Queue length: ", myQueue.length)
	fmt.Println(myQueue.tail)
	fmt.Println("////+2");
	myQueue.enqueue(5)
	myQueue.enqueue("five")
	fmt.Println("Queue length: ", myQueue.length)
	fmt.Println(myQueue.tail)
	

}
