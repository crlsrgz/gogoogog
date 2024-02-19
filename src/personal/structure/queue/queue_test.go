package main

import (
	"reflect"
	"testing"
)

func qnodeInt(value int) QNode {
	return QNode{value, nil}

}
func qnodeString(value string) QNode {
	return QNode{value, nil}

}
func createQueue() Queue {

	myQueue := Queue{}
	myQueue.enqueue(4)
	myQueue.enqueue("four")
	myQueue.enqueue(5)
	myQueue.enqueue("five")

	return myQueue
}
func TestQueue(t *testing.T) {
	myQueue := createQueue()
	if myQueue.head.value != 3 {
		t.Logf("Queue head is %v, want 4 (int)", myQueue.head.value)
	}

	if myQueue.tail.value != 4 {
		t.Fatalf("Queue head is %v, want five(string))", myQueue.tail.value)
	}
	myQueue.dumpQueue()
}

func TestQNodes(t *testing.T) {
	var intVar int
	node := qnodeInt(10)
	if reflect.TypeOf(node.value) != reflect.TypeOf(intVar) {

		t.Logf("Node value is not an int, value: %v, learn about types in go", node.value)
	}
	if node.value != 10 {
		t.Fatalf("Node value is %v, want 10 type int", node.value)
	}

}
