package main

import "fmt"

var arrayUnsortedOne = []int{40, 5, 34, 12, 33, 35, 13, 38, 27, 7, 6, 28, 44, 18, 1, 36, 11, 29, 9, 26, 20, 21, 15, 8, 25}

func shellSort(array []int, length int) {
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			j := i
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap)
				j = j - gap
			}
		}
	}

	fmt.Println(array)

}
func twoShellSort(array []int, length int) {
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			var j = i
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap)
				j = j - gap
			}
		}
	}
}

func swap(array []int, a int, b int) {
	array[a] = array[a] + array[b]
	array[b] = array[a] - array[b]
	array[a] = array[a] - array[b]
}

func main() {

	fmt.Println(arrayUnsortedOne)
	array := arrayUnsortedOne

	length := len(array)
	shellSort(array, length)
	twoShellSort(array, length)
	fmt.Println(array)

}
