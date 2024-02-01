package main

import "fmt"

func sum(lhs, rhs int) int {
	return lhs + rhs
}

func returnMultipleValues() (int, int, int) {
	return 1, 2, 3
}

func mixedTypes(animal string, number int) (string, int, string) {
	return "I have", number, animal
}

func recursive(n int) int {
	fmt.Println("recursion value:", n)
	if n == 1 {
		return 1
	}
	return n + recursive(n-1)
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(returnMultipleValues())
	fmt.Println(mixedTypes("dogs", 3))
	fmt.Println(recursive(5))
}
