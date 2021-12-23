package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in GO")
	content := "This needs to go inside file -> MYLEARNING"
	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./myfile.txt")

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is\n", string(dataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
