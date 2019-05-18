// main package hs examples shown
// in Hands-On Data Structures and Algorithms woth Golang
package main

// import fmt and container list package
import (
	"fmt"
	"container/list"
)

// main method
func main(){
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element= element.Next() {
		fmt.Println(element.Value.(int))
	}
}