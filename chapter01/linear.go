// main package has examples shown
// in Hans-On Data Structures and Algorithms with Golang
package main

import "fmt"

// import fmt package

/*
An algorithm is of linear complexity if the processing time or
stotage space is directly proportional to the number of input
elements to be processed O(n)
*/

// main method
func main() {
	var m [10]int
	var k int
	for k = 0; k < 10; k++ {
		m[k] = k * 200
		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
