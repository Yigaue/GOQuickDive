
// Go has only one looping construct, the for loop.

// The basic for loop has three components separated by semicolons:

package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	} 

	fmt.Print(sum , "\n")
}

// terminal: go run for.go