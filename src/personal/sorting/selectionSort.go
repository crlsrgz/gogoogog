package main

import "fmt"

var arrayUnsortedOne = []int{40, 5, 34, 12, 33, 35, 13, 38, 27, 7, 6, 28, 44, 18, 1, 36, 11, 29, 9, 26, 20, 21, 15, 8, 25}
var arrayUnsortedTwo = []int{15, 36, 1, 38, 17, 41, 22, 26, 30, 29, 44, 4, 20, 32, 3, 43, 39, 33, 12, 21, 42, 16, 31, 27, 28}

func selectionSort(array []int, length int) []int {
	var sortedArray []int = array
	var lowerIdx int

	for i := 0; i < length-1; i++ {
		// set the lower index to the first value
		lowerIdx = i
		//the lower value using the lower index
		var minVal = array[lowerIdx]
		for j := i; j < length-1; j++ {
			//if the minvalue is bigger than the neighbor chnage update the index
			if minVal > sortedArray[j+1] {
				minVal = sortedArray[j+1]
				lowerIdx = j + 1
			}

			fmt.Println(sortedArray)
		}

		fmt.Println(sortedArray)

		if i != lowerIdx {
			var temp = sortedArray[i]
			sortedArray[i] = sortedArray[lowerIdx]
			sortedArray[lowerIdx] = temp
		}
	}
	return sortedArray
}

func main() {
	length := len(arrayUnsortedOne)
	fmt.Println(arrayUnsortedOne)
	fmt.Println(selectionSort(arrayUnsortedOne, length))
	fmt.Println(selectionSort(arrayUnsortedTwo, length))
}
