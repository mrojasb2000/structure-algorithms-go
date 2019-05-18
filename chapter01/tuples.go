// main package hs examples shown
// in Hands-On Data Structures and Algorithms woth Golang
package main

import "fmt"

// import fmt and container list package

// gets the power series of integers a and returns tuple of square of a
// and cube of a
//func powerSeries(a int) (int, int) {
//	return a * a, a * a * a
//}

func powerSeries(a int) (square int, cube int) {
	square = a * a
	cube = a * a * a
	return
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, " Cube ", cube)
}
