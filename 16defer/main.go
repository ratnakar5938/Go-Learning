package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("Hello World") // executes at the end of the function
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Defer in GO")
	myDefer()

}

// Defer in GO
// 4 -> 3 -> 2 -> 1 -> 0
// Two -> One -> Hello World

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
