package main

import (
	"fmt"
	"math"
)

func power(base, exponent int) int {
	var result int = base
	for i := 0; i < exponent-1; i++ {
		result = result * base
	}
	return result

}

func multiplyPower(base, firstExp, secondExp int) int {
	var result int = base

	for i := 0; i < firstExp+secondExp-1; i++ {
		result = result * base
	}
	return result
}

func expontentGo(base, exponent int) int {
	operation := math.Pow(float64(base), float64(exponent))
	var result int = int(operation)
	return result
}
func main() {
	fmt.Printf("2^2 is equal to %v\n", power(2, 2))
	fmt.Printf("2^8 is equal to %v\n", power(2, 3))
	fmt.Printf("golang 2^8 is equal to %v\n", expontentGo(2, 3))
	fmt.Printf("2^3 * 2^4 is equal to %v\n", multiplyPower(2, 3, 4))
}
