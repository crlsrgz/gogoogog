package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (dog Dog) Speak() string {
	return "Guau!"
}

func (cat Cat) Speak() string {
	return "Miau"
}

func main() {
	var one Speaker = Dog{}
	var two Speaker = Cat{}

	fmt.Println(one.Speak())
	fmt.Println(two.Speak())
}
