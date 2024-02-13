package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return fmt.Sprintf("Going North")
	case South:
		return fmt.Sprintf("Going South")
	case East:
		return fmt.Sprintf("Going East")
	case West:
		return fmt.Sprintf("Going West")
	default:
		return "Other direction"
	}
}

func main() {
	north := North
	fmt.Println(north)
}
