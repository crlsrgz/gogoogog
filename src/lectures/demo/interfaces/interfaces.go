package main

import "fmt"

// explanation
type Resetter interface {
	Reset()
}
type Coordinate struct {
	x int
	y int
}
type Player struct {
	health   int
	position Coordinate
}

func (p *Player) Reset() {
	p.health = 100
	p.position = Coordinate{0, 0}
}

func Reset(r Resetter) {
	r.Reset()
}

// demo

type Preparer interface {
	PrepareDish()
}
type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}
func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}
func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("==Dish: %v==\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}
func main() {
	player := Player{50, Coordinate{10, 10}}
	fmt.Println(player)
	Reset(&player)
	fmt.Println(player)

	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Ceasar")}
	prepareDishes(dishes)
}
