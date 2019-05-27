package main

// main package has examples shown
// in Hands-On Data Structures and Algorithms with Golang

// import fmt package
import "fmt"

/*
An algorithm is of quadratic complexity if the
processing time is proportional to the square
of the number of input elements O(n^2).
*/

// main method
func main() {
	var k, l int
	for k = 1; k <= 10; k++ {
		fmt.Println(" Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			var x = l * k
			fmt.Println(x)
		}
	}
}
