package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your rating for our restaurant from 5 stars: ")
	reader := bufio.NewReader(os.Stdin)
	myRating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", myRating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		if numRating <= 5 {
			if numRating < 5 {
				numRating = numRating + 1
			}
			fmt.Println("Rating submitted,", numRating)
		} else {
			fmt.Println("Please enter between 1 and 5")
		}
	}

}
