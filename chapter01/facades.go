// main package has examples shown
// in Hands-On Data Structures and Algorithms with Golang

package main

// import fmt package
import (
	"fmt"
)

// Account struct
type Account struct {
	id          string
	accountType string
}

// Account class method create - creates account given AccountType
func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

//Account class method getById - get account given id string
func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by id")
	return account
}

// Account class method deleteId - deleteId given id string
func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}

// Customer struct
type Customer struct {
	name string
	id   int
}

// Customer class method create - create Customer given name
func (customer *Customer) create(name string) *Customer {
	customer.name = name
	return customer
}

// Transaction struct
type Transaction struct {
	id          string
	amount      float32
	srcAccount  string
	destAccount string
}

// Transaction class method create - create Transaction given src and dest account name and amount
func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("create transaction")
	transaction.srcAccount = srcAccountId
	transaction.destAccount = destAccountId
	transaction.amount = amount
	return transaction
}

// BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

// method NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

// BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

// main method
func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("21456", "87345", 1000)
	fmt.Println(transaction.amount)
}
