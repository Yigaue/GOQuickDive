//  Struct fields are accessed using a dot.

package main

import "fmt"
type Vertex struct {
	X int
	Y int
}

func main() {
	 v := Vertex{1, 2}
	 v.X = 4
	fmt.Println(v.X)
}

// terminal :  go run struct-fields.go