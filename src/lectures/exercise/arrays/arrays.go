//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import (
	"fmt"
	"math"
)

type Product struct {
	Name  string
	Price float64
}

func printProducts(list [4]Product) {
	var cost float64
	var totalItems int
	listLength := len(list)
	for i := 0; i < listLength; i++ {
		item := list[i]
		cost += item.Price

		if item.Name != "" {
			totalItems += 1
		}
	}
	fmt.Println("Last item:", list[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost:", math.Round(cost+100)/100)
}
func main() {
	shopingList := [4]Product{
		{"banana", 1.5},
		{"anana", 3.2},
		{"milk", 0.74},
	}
	printProducts(shopingList)
}
