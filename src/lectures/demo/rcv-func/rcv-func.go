package main

import "fmt"

type Coordinate struct {
	X, Y int
}

// normal function
func shiftBy(x, y int, coordinates *Coordinate) {
	coordinates.X += x
	coordinates.Y += y
}

// receiver function
// over wich 1. data struc, is operating on, 2. the function+params
func (coordinates *Coordinate) shiftBy(x, y int) {
	coordinates.X += x
	coordinates.Y += y
}

func main() {
	coord := Coordinate{2, 2}
	shiftBy(1, 1, &coord)
	fmt.Println("normal function:", coord)
	coord.shiftBy(2, 3)
	fmt.Println("receiver function:", coord)
}
