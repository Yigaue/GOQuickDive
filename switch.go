// A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "Linux":
		fmt.Println("Linux")
	default:
		// freesbd, opensbd,
		// plan9, windows...
		fmt.Printf("%s. \n", os)
	}

	// go run switch.go
}