package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"
	"fmt"
)

func main() {
	
	r := bufio.NewReader(os.Stdin)
	sum := 0
	for{
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}
		if inputErr == io.EOF{
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr);
			
		}
	}
	fmt.Printf("sum: %v\n", sum)
}
