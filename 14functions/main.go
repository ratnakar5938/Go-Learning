package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Functions in GO lang")
	superMan()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	myName := strings.TrimSpace(name)
	greeter(myName)

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	result, length, myMessage := proAdder(3, 5, 7, 9, 1)
	fmt.Println("Result is:", result)
	fmt.Println("Length of the slice", length)
	fmt.Println(myMessage)

}

func superMan() {
	fmt.Println("I am SUPERMAN")
}

func greeter(a string) {
	fmt.Printf("Hello %v from golang\n", a)
}

func adder(val1, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, int, string) {
	result := 0
	for _, value := range values {
		result += value
	}
	return result, len(values), "My message"
}
