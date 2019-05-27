// main package has examples shown
// in Hands-On Data Structures amd Algorithms with Golang
package main

import "fmt"

// import fmt package

/*
In the case of cubic complexity, the procesing
time of an algorithm is proportional to the
cube of the input elements O(n^3).
*/

// main method
func main() {
	var k, l, m int
	var arr [10][10][10]int
	for k = 0; k < 10; k++ {
		for l = 0; l < 10; l++ {
			for m = 0; m < 10; m++ {
				arr[k][l][m] = 1
				fmt.Println("Element value ", k, l, m, " is ", arr[k][l][m])
			}
		}
	}
}
