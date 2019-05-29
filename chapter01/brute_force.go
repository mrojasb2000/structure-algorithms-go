// main package has examples shown
// in Hands-On Data Structures and Algorithms with Golang
package main

import "fmt"

// findElement method given array and k element
func findElement(arr [10]int, k int) bool {
	var i int
	for i = 0; i < len(arr); i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

// main method
func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var check = findElement(arr, 10)
	fmt.Println(check)
	var check2 = findElement(arr, 9)
	fmt.Println(check2)
}
