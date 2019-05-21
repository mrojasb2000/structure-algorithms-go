// main package has examples shown
// in Hands-On Data Structures and Algot¡rithms with Golang
package main

// import fmt package
import (
	"fmt"
)

// DataTransferObjectFactory struct
type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

// DataTransferObjectFactory class method getDataTransferObject
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("New DTO of dtoType: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

// DataTransferObject interface
type DataTransferObject interface {
	getId() string
}

// Customer struct
type Customer struct {
	id   string // sequence generator
	name string
	ssn  string
}

// Customer class method getId
func (customer Customer) getId() string {
	fmt.Println("getting customer Id")
	return customer.id
}

// Employee struct {
type Employee struct {
	id   string
	name string
}

// Employee class method getId
func (employee Employee) getId() string {
	fmt.Println("getting employee Id")
	return employee.id
}

// Manager struct
type Manager struct {
	id   string
	name string
	dept string
}

// Manager class methd getId
func (manager Manager) getId() string {
	fmt.Println("getting manager Id")
	return manager.id
}

// Address struct
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

// Address class method getId
func (address Address) getId() string {
	fmt.Println("getting address Id")
	return address.id
}

// main method
func main() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}
	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer ", customer.getId())

	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee ", employee.getId())

	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager ", manager.getId())

	var address DataTransferObject = factory.getDataTransferObject("address")
	fmt.Println("Address ", address.getId())

}