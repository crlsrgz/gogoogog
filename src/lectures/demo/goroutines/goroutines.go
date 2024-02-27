package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {

	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capitalizeIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)

	}
	for i := 0; i < len(data); i++ {
		go capitalizeIt(data[i])
	}
	time.Sleep(20 * time.Millisecond)
	fmt.Printf("Capitalized: %c \n", capitalized)
	/* ═══  ═══ */
	fmt.Printf("\n\n/////////\n")
	var counter int = 0
	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)
		counter += 1
	}
	fmt.Println("Launch goroutines")
	fmt.Println(100 * time.Millisecond)

	go wait(100)
	go wait(900)
	go wait(1000)

	fmt.Println("launched counter=", counter)
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("waited 1100ms. counter=", counter)

}
