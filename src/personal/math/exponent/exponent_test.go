package main

import (
	"fmt"
	"testing"
)

//	func expontentGo(base, exponent float64) int {
//		operation := math.Pow(base, exponent)
//		var result int = int(operation)
//		return result
//	}
func TestExponential(t *testing.T) {
	firstBase := 2
	firstExp := 3
	var firstResult int = 9
	fmt.Println(firstBase, firstExp, firstResult)
	fmt.Println(power(firstBase, firstExp))
	if expontentGo(firstBase, firstExp) == power(firstBase, firstExp) {

		t.Logf("Ok")
	} else {
		t.Fatalf("not ok")
	}
	if power(firstBase, firstExp) != firstResult {
		t.Fatalf("not ok")
	}
}
