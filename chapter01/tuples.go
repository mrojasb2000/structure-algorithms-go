// main package hs examples shown
// in Hands-On Data Structures and Algorithms woth Golang
package main

import "fmt"

// import fmt and container list package

// gets the power series of integers a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, " Cube ", cube)
}
