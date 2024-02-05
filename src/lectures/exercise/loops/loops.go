//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//* Print integers 1 to 50, except:
	for i := 0; i < 51; i++ {
		//  - Print "Fizz" if the integer is divisible by 3
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("case FizzBuzz:", i)
		} else if i%3 == 0 {
			fmt.Println("case 3 Fizz:", i)
		} else if i%5 == 0 {
			fmt.Println("case 5 Fizz:", i)
		} else {
			fmt.Println(i)

		}

		//  - Print "Buzz" if the integer is divisible by 5
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
	}
}
