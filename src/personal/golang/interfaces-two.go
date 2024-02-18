package main

import "fmt"

type Animal interface {
	Speak() string
	Eat() string
}

type Bird struct{}
type Fish struct{}

func (b Bird) Speak() string {
	return "Pio"
}
func (f Fish) Speak() string {
	return "Blop"
}
func (f Fish) Eat() string {
	return "corn"
}
func (b Bird) Eat() string {
	return "seeds"
}

func main() {
	var one Animal = Bird{}
	var two Animal = Fish{}

	fmt.Println(one.Speak())
	fmt.Println(one.Eat())
	fmt.Println(two.Speak())
	fmt.Println(two.Eat())
}
