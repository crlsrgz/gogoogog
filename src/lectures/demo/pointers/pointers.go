package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// in this case the * next to teh variable is not needed.
	// the . after the variable does the trick
	counter.hits += 1
	fmt.Println("COunter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	value := 10
	// * for declaring a pointer type
	var valuePointer *int
	// & for assigning pointer/creating pointer
	valuePointer = &value
	valuePointerTwo := &value
	fmt.Println(value, valuePointer, valuePointerTwo)
	// * in this case will dereference the pointer,
	// and the operations on the dereferenced ponter will occur on the original data
	*valuePointer += 1

	fmt.Println(value, valuePointer, valuePointerTwo)

	// demo
	counter := Counter{}
	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)
	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

}
