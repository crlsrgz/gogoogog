package main

import "fmt"

type MyStructure struct {
	Name     string
	LastName string
	Age      int
}

func main() {

	person := MyStructure{"Ron", "weasley", 11}
	fmt.Println(person.Name)

}
