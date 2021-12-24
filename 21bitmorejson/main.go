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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "MyPlatform.com",
		"tags": ["web-dev", "js"]
	}
	`)

	var myCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for key, value := range myData {
		fmt.Printf("Key is %v, Value is %v and Type is %T\n", key, value, value)
	}
}
