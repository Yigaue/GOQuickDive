// factored import
package main 

import ( 
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite numebr is ", rand.Intn(10)) // N/B 'P' in Print is uppercase
}
// on terminal : go run package.go