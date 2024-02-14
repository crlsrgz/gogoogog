package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

// func (d Direction) String() string {
// 	switch d {
// 	case North:
// 		return fmt.Sprintf("Going North")
// 	case South:
// 		return fmt.Sprintf("Going South")
// 	case East:
// 		return fmt.Sprintf("Going East")
// 	case West:
// 		return fmt.Sprintf("Going West")
// 	default:
// 		return "Other direction"
// 	}
// }
/* 💡 Folowing pattern muss have the same order as the const list */
func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}
}
func main() {
	north := North
	fmt.Println(north)
}
