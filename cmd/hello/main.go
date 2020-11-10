package main

import (
	"fmt"
)

func main() {
	text := fmt.Sprintf("%s %s", "Hello", "World")

	fmt.Println(text)
}

// Exercise 5 Minute:
// Change to code above to concate 2 strings, load them into a variable
// and print the variable contents.
// Hint: Look at the documentation on the fmt package at https://pkg.go.dev/fmt#Sprintf
