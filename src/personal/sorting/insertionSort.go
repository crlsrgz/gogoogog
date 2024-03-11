package main

import "fmt"

var arrayUnsortedOne = []int{40, 5, 34, 12, 33, 35, 13, 38, 27, 7, 6, 28, 44, 18, 1, 36, 11, 29, 9, 26, 20, 21, 15, 8, 25}

func selectionSortOne(arr []int) []int {
	array := arr
	arrayLength := len(array)

	for i := 1; i < arrayLength; i++ {
		current := array[i]
		position := i - 1
	While:
		for {
			if position < 0 {
				break While
			}
			if array[position] > current {
				array[position+1] = array[position]
				position = position - 1
			} else {
				break While
			}

		}
		array[position+1] = current
	}
	return array
}
func selectionSortTwo(arr []int) []int {

	array := arr
	arrayLength := len(array)

	for i := 1; i < arrayLength; i++ {
		for j := i; j > 0 && array[j-1] > array[j]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}

	return arr
}
func main() {

	fmt.Println("unsorted array:", arrayUnsortedOne)
	fmt.Println("selection sort one:", selectionSortOne(arrayUnsortedOne))
	fmt.Println("selection sort one:", selectionSortTwo(arrayUnsortedOne))
}
