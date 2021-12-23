package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {

	fmt.Println("Welcome to web requests in GO")

	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Type of response is %T\n", response)
	defer response.Body.Close() // it's important to close the response

	dataBytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	content := string(dataBytes)
	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
