/* example of generics in golang */
package main

import (
	"fmt"
	"reflect"
)

func PrintSomething[T any, K any](entryOne T, entryTwo K) {
	fmt.Println(reflect.TypeOf(entryOne))
	fmt.Println(reflect.TypeOf(entryTwo))

	fmt.Println(entryOne, "<::::>", entryTwo)
}

type Insert interface {
	int | string | float64
}
type SmallThingOne[I Insert] struct {
	value  I
	number int
}

type SmallThingTwo[T any] struct {
	value  T
	number int
}

type ThreeRandomValues[K any, L any, I Insert] struct {
	first  K
	second L
	third  I
}

func main() {
	fmt.Println("//Go Generics//")
	PrintSomething(2, "thow")
	PrintSomething("string", 55)
	PrintSomething(1.36767, "floating")
	fmt.Printf("\n////\n\n")

	myThing := SmallThingTwo[string]{value: "hello", number: 22}
	fmt.Println(myThing)
	fmt.Printf("\n////\n\n")

	myThingOne := SmallThingOne[int]{value: 11, number: 33}
	fmt.Println(myThingOne)

	fmt.Printf("\n////\n\n")

	threeValues := ThreeRandomValues[string, int, float64]{"one", 2, 3.3333}
	fmt.Println(threeValues)

}
