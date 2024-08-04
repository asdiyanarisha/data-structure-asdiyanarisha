package main

import "fmt"

type Account struct {
	id          string
	accountType string
}

type Customer struct {
	name string
	id   int
}

type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by id")
	return account
}

func (account *Account) deleteById(id string) *Account {
	fmt.Println("delete account by id")
	return account
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

// BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId string,
	destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

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
