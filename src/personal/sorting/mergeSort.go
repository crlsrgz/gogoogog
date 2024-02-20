package main

import "fmt"

var unsArray = []int{5, 6, 8, 4, 2, 7, 1, 3}

func merge(left, right []int) []int {
	var arr = []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}

	}
	for ; i < len(left); i++ {
		arr = append(arr, left[i])
	}
	for ; j < len(right); j++ {
		arr = append(arr, right[j])
	}
	return arr
}
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// fmt.Println(arr)
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	// fmt.Println(first, second)
	fmt.Println(arr)
	return merge(first, second)
}
func main() {

	fmt.Println(unsArray)
	sorted := mergeSort(unsArray)
	fmt.Println(sorted)
}
