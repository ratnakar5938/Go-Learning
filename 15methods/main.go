package main

import (
	"fmt"
)

func main() {

	fmt.Println("Methods in Golang")
	ratnakar := User{"Ratnakar", "ratnakar@go.dev", true, 24}
	fmt.Printf("Name is %v and Email is %v\n", ratnakar.Name, ratnakar.Email)
	ratnakar.GetStatus()
	ratnakar.NewMail()
	fmt.Printf("Email is changed to: %v\n", ratnakar.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u *User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is:", u.Email)
}
