// The expression T(v) converts the value v to the type T.

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// terminal: go run type-conversions.go