package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "Hello"
	fmt.Println("Hello, World")
	fmt.Println("Happy", Pi, "Day")
}
// terminal: go run constants.go