package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("More on JSON in GO")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJs Bootcamp", 299, "MyPlatform.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 399, "MyPlatform.com", "bcd123", []string{"full-stack", "js"}},
		{"PRO Backend Developer", 599, "MyPlatform.com", "rat123", nil},
	}

	// package this data as json data
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
