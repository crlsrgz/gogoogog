package main

import (
	// "coursecontent/lectures/demo/pkg/msg"
	"fmt"
	"pkgLesson/display"
	"pkgLesson/msg"
)

func main() {
	fmt.Println("hello")
	// msg.Hi()
	// msg.Exciting("an Exciting message")
	msg.Exciting("Exciting message from mod in pkg")
	msg.Hi()
	display.Display("display from mod in pkg")

}
