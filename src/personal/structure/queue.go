package main

import "fmt"

type QueueNode struct {
	value int
	next  *QueueNode
}
type Queue struct {
	head   *QueueNode
	tail   *QueueNode
	length int
}

func createQueue() Queue {
	var queue = new(Queue)
	return *queue
}

func (queue *Queue) addToQueue(nodeValue int) *Queue {
	node := QueueNode{nodeValue, nil}

	queue.length++
	if queue.tail == nil {
		queue.head = &node
		queue.tail = queue.head
	}
	queue.tail.next = &node
	queue.tail = &node
	return queue
}
func (queue *Queue) removeFromQueue() int {
	if queue.head == nil {
		fmt.Println("error, the queue is empty")
		return -1
	}
	queue.length--
	newHead := queue.head
	queue.head = queue.head.next
	/* 💡 Automatic garbage collected? */
	newHead.next = nil

	return newHead.value
}
func (queue *Queue) peekQueue() int {
	if queue.head == nil {
		fmt.Println("error, the queue is empty, no head value")
		return -1
	}
	fmt.Println("peeked:", queue.head.value)
	return queue.head.value
}
func (queue *Queue) lengthQueue() int {
	if queue.head == nil {
		fmt.Println("error, no length, the queue is empty")
		return -1
	}
	fmt.Println("length:", queue.length)
	return queue.length
}

func main() {
	myQueue := createQueue()
	fmt.Println(myQueue.head)
	myQueue.removeFromQueue()
	myQueue.lengthQueue()

	myQueue.addToQueue(11)
	myQueue.addToQueue(22)
	myQueue.addToQueue(33)

	myQueue.lengthQueue()
	fmt.Println("head:", myQueue.head)
	fmt.Println("tail:", myQueue.tail)
	myQueue.removeFromQueue()
	myQueue.lengthQueue()
	myQueue.peekQueue()
	fmt.Println("head:", myQueue.head)
	fmt.Println("tail:", myQueue.tail)
	myQueue.lengthQueue()

}
