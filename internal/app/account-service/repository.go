package main

import (
	accountpb "github.com/chriswilliams1977/moneytree-protos/account"
	"log"
)

//interface to define data a repo can handle
type repository interface {
	Create(*accountpb.Account) (*accountpb.Account, error)
	GetAll() ([]*accountpb.Account, error)
	GetById(id string) (*accountpb.Account, error)
}

// Repository ...
type Repository struct {
	account *accountpb.Account
	accounts []*accountpb.Account
}

// Create a new account
func (repo *Repository) Create(account *accountpb.Account) (*accountpb.Account, error) {
	updated := append(repo.accounts, account)
	repo.accounts = updated
	return account, nil
}

// Get list of account based on customer,
// if customer name matches account,
// then return that account.
func (repo *Repository) GetAll() ([]*accountpb.Account, error){
	return repo.accounts, nil
}

func (repo *Repository) GetById(id string) (*accountpb.Account, error) {
	for _, account := range repo.accounts{
		if id == account.Number{
			log.Println("account is:", account)
			repo.account = account
		}
	}
	return repo.account, nil
}

