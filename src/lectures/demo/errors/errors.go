package main

import (
	"errors"
	"fmt"
)
type Stuff struct {
	values []int
}
func (stuff *Stuff) Get(index int) (int, error){
	if index > len(stuff.values){
		return 0, errors.New(fmt.Sprintf("no element at index %v", index))
	}else {
		return stuff.values[index], nil
	}
}
func main() {
	fmt.Println("//////")
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("value=", value)
	}
}
