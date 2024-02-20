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

func main() {
	fmt.Println("//////")
	fmt.Println(divide(10, 200))

}
