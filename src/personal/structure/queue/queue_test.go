package main 

import "testing"
import "reflect"

func qnodeInt(value int) QNode{
	return QNode{value, nil}

}
func qnodeString(value string) QNode{
	return QNode{value, nil}

}

func TestQNodes(t *testing.T){
	var intVar int
	node := qnodeInt(10) 
	if reflect.TypeOf(node.value) != reflect.TypeOf(intVar){

		t.Logf("Node value is not an int, value: %v, learn about types in go", node.value)
	}
	if node.value != 10 {
		t.Fatalf("Node value is %v, want 10 type int", node.value)
	}
	
}