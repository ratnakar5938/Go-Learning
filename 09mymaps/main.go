package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in GO lang")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("List of all languagess:", languages)
	fmt.Printf("JS is short for %v\n", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

}
