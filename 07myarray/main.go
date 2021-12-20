package main

import "fmt"

func main() {

	fmt.Println("Welcome to arrays in GO lang")

	var fruitList [4]string
	fruitList[0] = "Mango"
	// fruitList[1] = "Apple"
	fruitList[2] = "Papaya"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie list: ", vegList)
	fmt.Println("Length list is: ", len(vegList))

}
