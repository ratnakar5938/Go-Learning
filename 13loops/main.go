package main

import (
	"fmt"
)

func main() {

	fmt.Println("Loops in GO lang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			// break // ends the loop
			// rougueValue++
			// continue // skips the iteration
			goto gotolabel // jumps to the label when execution reaches this
		}
		fmt.Println("Value is", rougueValue)
		rougueValue++
	}

gotolabel:
	fmt.Println("Comes here when goto is used")

}
