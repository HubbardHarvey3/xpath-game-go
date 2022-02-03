package main

import (
	"fmt"
	"strings"
)

var path string

func main() {
	fmt.Println("Welcome to Xpath!  Choose your path: Easy, Medium or Hard")
	fmt.Print("Type your option: ")
	fmt.Scanln(&path)

	path = strings.ToLower(path)

	if path == "easy" || path == "medium" || path == "hard" {
		fmt.Printf("Great, you chose %v\n", path)
	} else {
		fmt.Println("So you have chosen death...")
	}

}
