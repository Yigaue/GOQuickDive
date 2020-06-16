// Variables declared without an explicit initial value are given their zero value.
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}

// terminal: go run zero.go