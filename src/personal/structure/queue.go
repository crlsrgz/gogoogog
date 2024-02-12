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

func addToQueue(queue *Queue, nodeValue int) *Queue {
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
func removeFromQueue(queue *Queue) int {
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
func peekQueue(queue *Queue) int {
	if queue.head == nil {
		fmt.Println("error, the queue is empty, no head value")
		return -1
	}
	fmt.Println("peeked:", queue.head.value)
	return queue.head.value
}
func lengthQueue(queue *Queue) int {
	if queue.head == nil {
		fmt.Println("error, mo length, the queue is empty")
		return -1
	}
	fmt.Println("length:", queue.length)
	return queue.length
}

func main() {
	myQueue := createQueue()
	fmt.Println(myQueue.head)
	removeFromQueue(&myQueue)
	lengthQueue(&myQueue)

	addToQueue(&myQueue, 11)
	addToQueue(&myQueue, 22)
	addToQueue(&myQueue, 33)

	lengthQueue(&myQueue)
	fmt.Println("head:", myQueue.head)
	fmt.Println("tail:", myQueue.tail)
	removeFromQueue(&myQueue)
	lengthQueue(&myQueue)
	peekQueue(&myQueue)
	fmt.Println("head:", myQueue.head)
	fmt.Println("tail:", myQueue.tail)
	lengthQueue(&myQueue)

}
