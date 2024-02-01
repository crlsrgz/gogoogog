//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var myColor = "red"
	fmt.Println(myColor)
	birth, age := 1924, 100

	fmt.Println("Birth and age:", birth, "/", age)
	var (
		name     = "John"
		lastName = "Jones"
	)
	fmt.Println("My name is:", name, lastName)

	var ageInDays int
	fmt.Println("ageInDays", ageInDays)
	ageInDays = 100 * 365
	fmt.Println("ageInDays", ageInDays)

}
