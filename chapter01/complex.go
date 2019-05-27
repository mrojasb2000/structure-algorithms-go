// main package has examples shown
// in Hands-On Data Structure and Algorithms with Golang
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	var m [10]int
	var k int
	for k = 0; k < len(m); k++ {
		m[k] = k + 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
