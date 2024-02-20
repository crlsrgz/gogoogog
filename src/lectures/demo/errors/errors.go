package main

import (
	"errors"
	"fmt"
)

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return rhs / lhs, nil
	}
}
type DivError struct {
	a, b int
}
func (d *DivError) Error() string {
	return fmt.Sprintf("Cannot divide by Zero:%d /%d", d.a, d.b)
}
func div(a, b int) (int, error){
	if b == 0 {
		return 0, &DivError{a, b}
	} else {
		return a/b, nil
	}
}
func main() {
	fmt.Println("//////")
	fmt.Println(divide(10, 200))
	fmt.Println(divide(10, 0))
	fmt.Println("//////")

	answerOne, err := div(9,0)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println("div answer is", answerOne)
}
