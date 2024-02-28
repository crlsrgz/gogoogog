package main

import (
	"fmt"
	"testing"
)

func createSSL(arrAdd, arrRemove []int) []int {
	firstList := List{}
	for i := 0; i < len(arrAdd); i++ {
		firstList.addNode(arrAdd[i])
	}
	for i := 0; i < len(arrRemove); i++ {
		firstList.removeByValue(arrRemove[i])
	}
	arr := firstList.printValues()
	return arr

}

var firstTest = [][]int{
	{2, 4, 6}, // add
	{3},       // remove
	{2, 4, 6}, // check
}

var secondTest = [][]int{
	{5, 10, 15, 16, 18, 20, 25, 30, 35, 37, 40}, // add
	{16, 18, 37},                    // remove
	{5, 10, 15, 20, 25, 30, 35, 40}, // check
}

func TestOne(t *testing.T) {
	listValues := createSSL(firstTest[0], firstTest[1])
	for i := 0; i < len(firstTest[2]); i++ {
		if listValues[i] != firstTest[2][i] {
			t.Fatalf("value at index %v = %v, want %v", i, listValues[i], firstTest[2][i])

		}

	}
	fmt.Printf("result: %v | %v :expected\n\n", listValues, firstTest[2])
}

func TestTwo(t *testing.T) {
	listValues := createSSL(secondTest[0], secondTest[1])
	for i := 0; i < len(secondTest[2]); i++ {
		if listValues[i] != secondTest[2][i] {
			t.Fatalf("value at index %v = %v, want %v", i, listValues[i], secondTest[2][i])

		}

	}
	fmt.Printf("result: %v | %v :expected\n\n", listValues, secondTest[2])
}
