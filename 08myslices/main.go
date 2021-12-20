package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Slices in GO lang")
	var fruitList = []string{"Apple", "Papaya", "Peach"}
	fmt.Printf("Type of Fruit List is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Printf("My fruit list is: %v\n", fruitList)

	// fruitList = append(fruitList[1:])
	// fruitList = append(fruitList[1:3])
	// fruitList = append(fruitList[1:-1]) // this works in other languages but doesn't work in GOlang
	fmt.Printf("My fruit list is: %v\n", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 988
	highScores[2] = 478
	highScores[3] = 788

	highScores = append(highScores, 455, 888)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
