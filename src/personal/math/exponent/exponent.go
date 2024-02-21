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
func powerRecursive(base, exponent int)int {
	if exponent == 0 {
		return 1
	}	
	fmt.Printf("rec val %vx%v= %v\n",base, (powerRecursive(base, exponent-1)),(base * powerRecursive(base, exponent-1)))
	return base * powerRecursive(base, exponent-1)
}

func multiplyPower(base, firstExp, secondExp int) int {
	var result int = base

	for i := 0; i < firstExp+secondExp-1; i++ {
		result = result * base
	}
	return result
}
func exponential(n int) int{
	if n == 0 {
		return 1
	}
	return n * exponential(n-1)
}

func expontentGo(base, exponent int) int {
	operation := math.Pow(float64(base), float64(exponent))
	var result int = int(operation)
	return result
}
func main() {
	fmt.Printf("2^2 is equal to %v\n", power(2, 2))
	fmt.Printf("2^8 is equal to %v\n", power(2, 3))
	fmt.Printf("3^4 is equal to %v\n", power(3, 4))
	fmt.Printf("2^8 is equal to %v, using recursion\n", powerRecursive(2, 3))
	fmt.Printf("golang 2^8 is equal to %v\n", expontentGo(2, 3))
	fmt.Printf("2^3 * 2^4 is equal to %v\n", multiplyPower(2, 3, 4))
	fmt.Printf("Exponential of 4 is %v\n", exponential(4))
	fmt.Printf("Exponential of 5 is %v\n", exponential(5))
}
