package main

import (
	"fmt"
)

func main() {

	fmt.Println("Structs in GO lang")
	// no inheritance in golang : No super or parent

	hitesh := User{"Hitesh", "hitesh@dev.go", true, 25}
	fmt.Printf("User details are : %+v\n", hitesh)
	fmt.Printf("Name is %v\nEmail is %v\n", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
