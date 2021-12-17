package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating for our restraurant: ")

	// comma ok || error ok syntax
	myRating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", myRating)
	fmt.Printf("Types of this rating is %T", myRating)

}
