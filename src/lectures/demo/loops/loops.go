package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("/////")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Even Number:", i)
	}

	//while
	myLoop := 0

	for myLoop < 6 {

		fmt.Println("While with for: ", myLoop)
		myLoop++
	}
	fmt.Println("//////")
	mySecondLoop := 0

	for mySecondLoop < 6 {

		mySecondLoop++
		fmt.Println("While with for: ", mySecondLoop)
	}
	fmt.Println("//////")

}
