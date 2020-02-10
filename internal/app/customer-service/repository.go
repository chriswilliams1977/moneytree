package main

import (
	// Import the generated protobuf code
	customerpb "github.com/chriswilliams1977/moneytree-protos/customer"
	accountpb "github.com/chriswilliams1977/moneytree-protos/account"
	"log"
)

//interface to define data a repo can handle
type repository interface {

	//takes a customer a returns and customer
	Create(*customerpb.Customer) (*customerpb.Customer, error)
	//return slice of customers
	GetAll() ([]*customerpb.Customer, error)
	GetById(id string) (*customerpb.Customer, error)
	CreateAccount(*accountpb.Account) (*accountpb.Account, error)
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
//Create repo object and implement interface to match type
type Repository struct {
	customer *customerpb.Customer
	customers []*customerpb.Customer
	accounts []*accountpb.Account
}

// Create a new customer
func (repo *Repository) Create(customer *customerpb.Customer) (*customerpb.Customer, error) {
	updated := append(repo.customers, customer)
	repo.customers = updated
	return customer, nil
}

// GetAll consignments
func (repo *Repository) GetAll() ([]*customerpb.Customer, error) {
	return repo.customers, nil
}

func (repo *Repository) GetById(id string) (*customerpb.Customer, error) {
	for _, customer := range repo.customers{
		if id == customer.AccountNumber {
			log.Println("customer is:", customer)
			repo.customer = customer
		}
	}
	return repo.customer, nil
}

//Create a new account record
func (repo *Repository) CreateAccount(account *accountpb.Account) (*accountpb.Account, error) {
	updated := append(repo.accounts, account)
	repo.accounts = updated
	return account, nil
}
