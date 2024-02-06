package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func input() {
	var input string
	var number int

	for {
		fmt.Print("integer: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input, integer:")
			continue
		}

		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not an integer, integer:")
			continue
		}

		// Successful conversion, exit the loop
		break
	}

	fmt.Printf(" Integer = %d\n", number)
}

func myArrays() {
	input()
	fmt.Println("enter the first value of the array")
	var firstValue int

	var checkFirstValue int = 1

	fmt.Scanln(&firstValue)
	fmt.Println("FirstValue:", firstValue)

	fmt.Println("checkFirstValue: ", reflect.TypeOf(firstValue) == reflect.TypeOf(checkFirstValue))

	if reflect.TypeOf(firstValue) != reflect.TypeOf(checkFirstValue) {
		fmt.Println("value should be number")
		fmt.Scanln(&firstValue)
	}

	for firstValue >= 10 {
		fmt.Println("value should be less than 10")
		fmt.Scanln(&firstValue)
	}

	var myFirstArray [3]int
	myFirstArray[0] = firstValue

	fmt.Println(myFirstArray)

	mySecondArray := [3]int{6, 0, 0}
	fmt.Println(mySecondArray)

	myThirdArray := [...]int{7, 0, 9}
	fmt.Println(myThirdArray)

	myFourthArray := [4]int{8, 0, 10}
	fmt.Println(myFourthArray)

}

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		/* 💡 Always make  a copy */
		room := rooms[i] //assigned to variable, concurrency
		if room.cleaned {
			fmt.Println(room.name, "is clean")

		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	myArrays()

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Kitchen"},
		{name: "Parking"},
	}

	rooms[1].cleaned = true

	checkCleanliness(rooms)
}
