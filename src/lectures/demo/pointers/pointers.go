package main

import "fmt"

func main() {
	value := 10
	// * for declaring a pointer type
	var valuePointer *int
	// & for assigning pointer/creating pointer
	valuePointer = &value
	valuePointerTwo := &value
	fmt.Println(value, valuePointer, valuePointerTwo)
	// * in this case will dereference the pointer,
	// and the operations on the dereferenced ponter will occur on the original data
	*valuePointer += 1
	fmt.Println(value, valuePointer, valuePointerTwo)
}
