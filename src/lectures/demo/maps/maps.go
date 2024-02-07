package main

import "fmt"

func main() {
	// Create
	myFirstMap := make(map[string]int)
	mySecondMap := map[string]int{
		"item 1": 1,
		"item 2": 2,
		"item 3": 3,
	}
	// Insert
	myFirstMap["item 1"] = 11
	mySecondMap["item 4"] = 4
	// Read
	var itemOne int = myFirstMap["item 1"]
	var itemTen int = myFirstMap["item 10"] // returns an empty variable if it doesn't exist
	fmt.Println(itemOne, itemTen)
	// Check existence
	item, found := mySecondMap["item 10"]
	fmt.Println("assigned unexisting valu to variable", item)
	if !found {
		fmt.Println("Not found")
	}

	fmt.Println(myFirstMap)

	fmt.Println(mySecondMap)
	// Delete
	delete(mySecondMap, "item 1")
	fmt.Println(mySecondMap)
	// Iterate
	for key, value := range mySecondMap {
		fmt.Print("Iterate_ ", key, ":", value, "\n")
	}
}
