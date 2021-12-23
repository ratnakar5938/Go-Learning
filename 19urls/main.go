package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=iuhgfisgfd"

func main() {

	fmt.Println("Urls in GO")
	fmt.Println(myurl)

	// parsing
	result, err := url.Parse(myurl)
	checkNilErr(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=ratnakar",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
