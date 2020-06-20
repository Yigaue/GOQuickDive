// Slices are like references to array
// A slice does not store any data, it just describes a section of an underlying array.

package main

import "fmt" 

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "Apollos"
	fmt.Println(a, b)
	fmt.Println(names)

}

// terminal: go run slices-pointers.go