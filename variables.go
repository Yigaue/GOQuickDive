//The var statement declares a list of variables; as in function argument lists, the type is last.

//A var statement can be at package or function level. We see both in this example.

package main

import "fmt"

var name, gender, occupation string

func main () {
	var i int
	fmt.Println(i, name, gender, occupation)
}

// go run variables.go