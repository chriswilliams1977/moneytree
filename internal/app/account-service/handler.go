package main

import (
	"context"
	"log"

	accountpb "github.com/chriswilliams1977/moneytree-protos/account"
)

type handler struct {
	repository
}

// FindAvailable accounts
func (h *handler) GetAccounts(ctx context.Context, req *accountpb.Request, res *accountpb.Response) error {

	accounts, _ := h.repository.GetAll()
	res.Accounts = accounts
	return nil

}

func (h *handler) GetAccountByNumber(ctx context.Context, req *accountpb.Request, res *accountpb.Response) error {

	log.Println("Calling GetCustomerById")

	account, _ := h.repository.GetById(req.AccountNumber)
	res.Account = account

	return nil
}

// Create a new account
func (h *handler) CreateAccount(ctx context.Context, req *accountpb.Account, res *accountpb.Response) error {
	// Save our customer
	account, err := h.repository.Create(req)
	if err != nil {
		return err
	}

	res.Account = account
	return nil
}



