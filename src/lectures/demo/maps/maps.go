// Maps are unordere, iteration with a for loop might yield different order
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
	// Iterate without keys
	for _, value := range mySecondMap {
		fmt.Print("Iterate-no-keys_ :", value, "\n")

	}
	fmt.Println("===================")
	shopingList := make(map[string]int)
	shopingList["eggs"] = 11

	fmt.Println(shopingList)

	_, foundMilk := shopingList["milk"]
	if !foundMilk {
		fmt.Println("no milk found with an empty variable")
	}
	shopingList["milk"] = 5
	fmt.Println(shopingList)
	if !foundMilk {
		fmt.Println("no milk found with an empty variable")
	}
	_, foundMilk = shopingList["milk"]
	if !foundMilk {
		fmt.Println("no milk found with an empty variable")
	} else {

		fmt.Println("milk found with an empty variable")
	}
}
