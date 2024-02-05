package main

import (
	"fmt"
)

func recursionSimple(number int) int {
	fmt.Println("Value:", number)

	if number == 1 {
		return 1
	}
	return number + recursionSimple(number-1)
}

func countDown(numb int) int {
	fmt.Println("number is:", numb)
	if numb == 0 {
		return 1
	}
	newNumber := numb - 1

	countDown(newNumber)

	return numb
}

func main() {
	fmt.Println("recursionSimple", recursionSimple((6)))
	fmt.Println("countDown", countDown(6))
}
