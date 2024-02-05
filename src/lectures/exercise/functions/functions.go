//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func personName(name string) string {
	return name
}

func textMessage(textMessage string) string {
	return textMessage
}

func addNumbers(a, b, c int) int {
	return a + b + c
}

func returnNumber(n int) int {
	return n
}

func returnTwoNumbers(a, b int) (int, int) {
	return a, b
}
func returnOtherTwoNumbers() (int, int) {
	return 2, 7
}
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func main() {

	fmt.Println("personName: ", personName("Peter"))
	fmt.Println("textMessage: ", textMessage("Good day"))
	fmt.Println("addNumbers: ", addNumbers(1, 2, 4))
	fmt.Println("returnNumber: ", returnNumber(7))
	fmt.Println(returnTwoNumbers(5, 7))
	fmt.Println("addThreeNumbers: ", addThreeNumbers(returnNumber(2), 3, 5))

}
