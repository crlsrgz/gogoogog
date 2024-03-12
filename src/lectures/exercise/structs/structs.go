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

import "fmt"

type Len interface {
	int | float32
}

type Rectangle[L Len] struct {
	len L
	wid L
}

func area(len, wid int) int {
	return len * wid
}

func perimeter(len, wid int) int {
	return (len * 2) + (wid * 2)
}

func areaTwo(len, wid int) int {
	return len * wid
}

func perimeterTwo(len, wid int) int {
	return (len * 2) + (wid * 2)
}
func main() {
	recOne := Rectangle[int]{3, 4}

	fmt.Println(area(recOne.len, recOne.wid))
	fmt.Println(perimeter(recOne.len, recOne.wid))

}
