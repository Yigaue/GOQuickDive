//A var declaration can include initializers, one per variable.

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

package main

import "fmt"

var i, name, gender, occupation = 2, "Sonia", "male", "Developer"

func main(){
	fmt.Println(1, name, gender, occupation)
}

// on terminal: go run variables-with-initialisers.go