//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"strconv"
)

func input(valueName string) int {
	var input string
	var number int

	for {
		fmt.Print("Whole number for ", valueName, ": ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input, Whole number for ", valueName, ": ")
			continue
		}

		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not a Whole number, Whole number for ", valueName, ": ")
			continue
		}

		// Successful conversion, exit the loop
		break
	}
	return number
	// fmt.Printf(" Integer = %d\n", number)
}

type Rectangle struct {
	Width     int
	Length    int
	Perimeter int
	Area      int
}

func main() {
	widthValue := input("width")
	lengthValue := input("length")
	//fmt.Println(widthValue, lengthValue)

	rectangleOne := Rectangle{}
	rectangleOne.Width = widthValue
	rectangleOne.Length = lengthValue

	rectangleArea := rectangleOne.Width * rectangleOne.Length
	rectanglePerimeter := (rectangleOne.Width * 2) + (rectangleOne.Length * 2)
	rectangleOne.Area = rectangleArea
	rectangleOne.Perimeter = rectanglePerimeter
	fmt.Println(rectangleOne, " Perimeter: ", rectanglePerimeter, "; ", "Area: ", rectangleArea)

}
