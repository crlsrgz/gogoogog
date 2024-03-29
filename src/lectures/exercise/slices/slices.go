//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printParts(parts []Part) {

	fmt.Println("\n----Parts----")
	for i := 0; i < len(parts); i++ {

		element := parts[i]
		fmt.Println(i, ".- ", element)

	}
}

func main() {
	//partsExample := make([]Part, 3)
	parts := []Part{"cog", "cable", "fuse"}
	printParts(parts)

	parts = append(parts, "led")
	parts = append(parts, "wheel")
	printParts(parts)

	parts = parts[3:]
	printParts(parts)
}
