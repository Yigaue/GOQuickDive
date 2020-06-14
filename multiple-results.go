package main

import "fmt"

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("Keep", "Learning")
	fmt.Println(a, b)
} 

// go run multiple-results.go