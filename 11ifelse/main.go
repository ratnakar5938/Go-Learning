package main

import (
	"fmt"
)

func main() {

	fmt.Println("If else in go lang")

	loginCount := 24
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login counts"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Printf("%v is less than 10\n", num)
	} else {
		fmt.Printf("%v is greater than 10\n", num)
	}

}
