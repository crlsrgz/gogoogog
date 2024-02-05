//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func checkAge(stage int) string {
	// var age int = stage
	switch age := stage; {
	case age == 420:
		fmt.Println("next case is joke for a fallthrough")
		fallthrough
	case age > 0 && age < 1:
		return "newborn"
	case age >= 1 && age < 4:
		return "toddler"
	case age >= 4 && age < 13:
		return "child"
	case age >= 13 && age < 18:
		return "teenager"
	default:
		return "adult"

	}
}

func main() {
	fmt.Println("age is 0", checkAge(0))
	fmt.Println("age is 1", checkAge(1))
	fmt.Println("age is 2", checkAge(2))
	fmt.Println("age is 3", checkAge(3))
	fmt.Println("age is 4", checkAge(4))
	fmt.Println("age is 5", checkAge(5))
	fmt.Println("age is 6", checkAge(6))
	fmt.Println("age is 7", checkAge(7))
	fmt.Println("age is 8", checkAge(8))
	fmt.Println("age is 9", checkAge(9))
	fmt.Println("age is 10", checkAge(10))
	fmt.Println("age is 12", checkAge(12))
	fmt.Println("age is 15", checkAge(15))
	fmt.Println("age is 17", checkAge(17))
	fmt.Println("age is 18", checkAge(18))
	fmt.Println("age is 21", checkAge(21))
	fmt.Println("age is 420", checkAge(420))
	fmt.Println("age is 99", checkAge(99))
}
