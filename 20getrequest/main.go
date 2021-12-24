package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myUrl string = "http://localhost:8000/get"

func main() {

	fmt.Println("Get Requests in GO")
	PerformGetRequest(myUrl)

}

func PerformGetRequest(geturl string) {
	response, err := http.Get(geturl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(content)
	// fmt.Println(string(content))

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}
