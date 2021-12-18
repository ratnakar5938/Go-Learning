package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of GoLang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.July, 10, 22, 22, 0, 0, time.UTC)
	fmt.Println(createDate.Format(time.UnixDate))

}
