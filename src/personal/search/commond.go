package main

import "fmt"

func gCommonDen(x, y int) int {
	var tmp int
	for y != 0 {
		tmp = x
		x = y
		y = tmp % y
	}
	return x
}

func main() {
	fmt.Println(gCommonDen(20, 8))
	fmt.Println(gCommonDen(60, 96))
}
