// main package has examples shown
// in Hands-On Data Structure and Algorithms with Golang

package main

// importing fmt package
import (
	"fmt"
)

// IRealObject interface
type IRealObject interface {
	performAction()
}

// RealObject struct
type RealObject struct{}

// RealObject class implement IRealObject interface.
// The class has method performAction.
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

// VirtualProxy struct
type VirtualProxy struct {
	realObject *RealObject
}

// CirtualProxy class method performAction
func (virtuaProxy *VirtualProxy) performAction() {
	if virtuaProxy.realObject == nil {
		virtuaProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	virtuaProxy.realObject.performAction()
}

func main() {
	var object = VirtualProxy{}
	object.performAction()
}
