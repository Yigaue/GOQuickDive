// Go's return values may be named. If so, they are treated as variables defined at the top of the function.

package main

import "Fmt"

func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}