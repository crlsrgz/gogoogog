package main

import "fmt"

func main() {

	runeOne := 'A'
	runeTwo := 'B'

	fmt.Println("these are runes for A and B", runeOne, runeTwo)

	name := "The Name"
	fmt.Println(name)

	var exampleOne = 1
	var exampleTwo int = 2
	var exampleThree int

	var aOne, bOne, cOne = 1, 2, "Three"
	// Block assignment
	var (
		aTwo int    = 1
		bTwo int    = 2
		cTwo string = "Three"
	)

	fmt.Println(exampleOne)
	fmt.Println(exampleTwo)
	fmt.Println(exampleThree)
	fmt.Println(aOne, bOne, cOne)
	fmt.Println(aTwo, bTwo, cTwo)

	exampleThree = 3
	exampleFour := 4 /* 💡 Type inference with := */

	exampleFive := exampleFour
	fmt.Println(`exampleFour:`, exampleFour, " exampleFive: ", exampleFive)

	exampleFive = 5
	fmt.Println(`exampleFour:`, exampleFour, " exampleFive: ", exampleFive)
	fmt.Println(`exampleFour:`, exampleFour, " exampleFive: ", exampleFive)

	// Unassigned values
	var emptyString string
	var emptyNumber int
	fmt.Println("emptyString: ", emptyString, "emptyNumber", emptyNumber)

	const OneConstant = "Pascal Case for constants"
	const AnotherConstant = 3

	fmt.Println(OneConstant, AnotherConstant, "\n ///////////////////")

	// Ignore a variable
	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2) // hello world
}
