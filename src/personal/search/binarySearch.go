package main

import "fmt"

func binarySearch(array []int, length int, needle int) map[string]int {

	var lowBound int = 0
	var highBound int = length
	var mid int = 0

	var result = map[string]int{}
	result["position"] = -1

	for lowBound < highBound {
		// if lowBound >= highBound {
		// 	break
		// }
		mid = (lowBound + highBound) / 2

		if needle == array[mid] {
			result["position"] = mid
			//fmt.Println("hello result")
			return result
		} else if needle > array[mid] {
			lowBound = mid + 1
		} else if needle < array[mid] {
			highBound = mid
		}

	}
	return result
}

func binarySearchTwo(array []int, length int, needle int) map[string]int {
	var low = 0
	var high = length
	var mid = 0

	var result = map[string]int{}
	result["position"] = -1

	for low < high {
		mid = low + (high-low)/2
		if array[mid] < needle {
			low = mid + 1
		} else {
			high = mid
		}
	}
	result["position"] = low
	return result
}
func main() {
	sortedArray := []int{1, 5, 6, 7, 8, 9, 11, 12, 13, 15, 18, 20, 21, 25, 26, 27, 28, 29, 33, 34, 35, 36, 38, 40, 44}
	//fmt.Println(sortedArray)
	var length int = len(sortedArray)
	var needle int = 29
	fmt.Println(binarySearch(sortedArray, length, needle))
	fmt.Println(binarySearchTwo(sortedArray, length, needle))

}
