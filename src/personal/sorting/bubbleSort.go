package main

import "fmt"

var arrayUnsortedOne = []int{5, 6, 8, 4, 2, 7, 1, 3}

func main() {

	fmt.Println("Unsorted array: ", arrayUnsortedOne)

	fmt.Println(bubbleSort(arrayUnsortedOne))
}

func bubbleSort(array []int) []int {
	var length int = len(array)
	fmt.Println("len", length)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if array[j+1] < array[j] {
				var tmp int = array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
			fmt.Println("step:", i, " ", array)

		}

		fmt.Println("///")
	}
	return array
}
