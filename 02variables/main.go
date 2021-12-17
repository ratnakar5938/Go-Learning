package main

import "fmt"

// jwtToken := 300000 // these are not allowed outside the function body
var jwtToken = 300000

const LoginToken string = "gfaifiugi" // First letter capitalization make the variable available everywhere ie public

func main() {
	var username string = "Ratnakar"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T, value: %v\n", username, username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("variable is of type: %T, value: %v\n", isLogged, isLogged)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type: %T, value: %v\n", smallval, smallval)

	// var smallfloat float32 = 255.45575496745967954867549
	var smallfloat float64 = 255.45575496745967954867549 // provides more precision than float32
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T, value: %v\n", smallfloat, smallfloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable of type: %T\n", anotherVariable)

	// implicit type
	var website = "Learn Code Online"
	fmt.Println(website)

	// no var style
	numberOfUsers := jwtToken + 3000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable of type: %T\n", LoginToken)
}
