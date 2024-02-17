package display

import "fmt"

func Display(msg string) {
	fmt.Println(msg)
}

/* 💡 The following function can't be accessed by the main package beacause its name doesn't start with a capital letter */
func hello(msg string) {
	fmt.Println(msg)
}
