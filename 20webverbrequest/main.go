package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myGETurl string = "http://localhost:8000/get"

const myPOSTurl string = "http://localhost:8000/post"

const myPOSTFORMurl string = "http://localhost:8000/postform"

func main() {

	fmt.Println("Get Requests in GO")
	PerformGetRequest(myGETurl)
	PerformPostJsonRequest(myPOSTurl)
	PerformPostFormRequest(myPOSTFORMurl)

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

func PerformPostJsonRequest(postUrl string) {
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"RATNAKAR",
			"course":"Let's Go With GoLang",
			"price":0
		}
	`)

	response, err := http.Post(postUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest(postFormUrl string) {
	data := url.Values{}
	data.Add("firstname", "Ratnakar")
	data.Add("lastname", "Sahoo")
	data.Add("email", "Ratnakar@go.dev")
	data.Add("country", "India")

	response, err := http.PostForm(postFormUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
