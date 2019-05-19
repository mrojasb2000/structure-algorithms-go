// main package has examples shown
// in Hands-On Data Structures and Algorithms with Golang
package main

// import fmt package
import (
	"fmt"
)

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

// Leaflet class name perform
func (leaf *Leaflet) perfom() {
	fmt.Println("Leaflet " + leaf.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// Branch class name perfom
func (branch *Branch) perfom() {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perfom()
	}
	for _, branch := range branch.branches {
		branch.perfom()
	}
}

// Branch class method add Leaflet
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// Branch class method addBranch branch
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// Branch class method getLeaflet
func (branch *Branch) getLeaflet() []Leaflet {
	return branch.leafs
}

// main method
func main() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}

	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perfom()
}
