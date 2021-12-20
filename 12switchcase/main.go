package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch Case in GO")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Value of dice is: %v\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		// fallthrough // goes to the next case even if it is not true
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("Not going to happen")
	}

}
