// A struct is a collection of fields.

package main

import "fmt" 

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

// terminal: go run structs.go