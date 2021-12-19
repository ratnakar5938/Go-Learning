package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on poinetrs")

	// var ptr *int
	// fmt.Printf("value of pointer is %v\n", ptr)

	myNumber := 34
	var myptr *int = &myNumber
	fmt.Printf("The address of the variable is: %v\nValue of variable is: %v\n", myptr, *myptr)

	*myptr = *myptr * 2
	fmt.Println("New value of the number", myNumber)
}
