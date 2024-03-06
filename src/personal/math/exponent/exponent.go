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
func powerRecursive(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	fmt.Printf("rec val %vx%v= %v\n", base, (powerRecursive(base, exponent-1)), (base * powerRecursive(base, exponent-1)))
	return base * powerRecursive(base, exponent-1)
}

func multiplyPower(base, firstExp, secondExp int) int {
	var result int = base

	for i := 0; i < firstExp+secondExp-1; i++ {
		result = result * base
	}
	return result
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Multiply two expressions with the same base
// x^n * x^m = x^n+m

func multiplyTwoPowerSameBase(base, firstExponent, secondExponent int) int {
	var result int = base
	totalExponent := firstExponent + secondExponent

	for totalExponent > 1 {
		result = result * base
		fmt.Printf("check: %v\n", result)

		totalExponent -= 1
	}

	return result
}

// Multiply two expressions with the same base
// x^n / x^m = x^n-m

func divideTwoPowerSameBase(base, firstExponent, secondExponent int) int {
	var result int = base
	totalExponent := firstExponent - secondExponent
	for totalExponent > 1 {
		result = result * base
		totalExponent -= 1
	}
	return result
}

// Power a base with exponent
// (x^n)^m = x^n*m

func doublePower(base, firstExponent, secondExponent int) int {
	var result int = base
	totalExponent := firstExponent * secondExponent
	for totalExponent > 1 {
		result = result * base
		totalExponent -= 1
	}
	return result

}

// ///////////////////
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
	fmt.Printf("Exponential of 4 is %v\n", factorial(4))
	fmt.Printf("Exponential of 5 is %v\n", factorial(5))
	fmt.Printf("two power same base mult %v\n", multiplyTwoPowerSameBase(2, 3, 2))
	fmt.Printf("two power same base divide %v\n", divideTwoPowerSameBase(2, 4, 2))
	fmt.Printf("(2^3)^3 %v\n", doublePower(2, 3, 3))

}
