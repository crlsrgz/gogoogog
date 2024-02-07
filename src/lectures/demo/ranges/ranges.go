package main

import "fmt"

func main() {

	slice := []string{"hello", "world", "!"}
	for i, element := range slice {
		fmt.Println(i, element)

		for _, ch := range element {
			fmt.Printf("   %q\n", ch)
		}
	}
	fmt.Println("------")
	for _, element := range slice {
		fmt.Println(element)
		for _, character := range element {
			fmt.Printf("   %q\n", character)
		}
	}
}
