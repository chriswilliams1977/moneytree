package main

import (
	"context"
	"log"

	accountpb "github.com/chriswilliams1977/moneytree-protos/account"
	customerpb "github.com/chriswilliams1977/moneytree-protos/customer"
)

//The handler implements the custer server functions as per the interface contract
type handler struct {
	repository
	accountClient accountpb.AccountServiceClient
}

// CreateCustomer - we created just one method on our service,
func (h *handler) CreateCustomer(ctx context.Context, req *customerpb.Customer, res *customerpb.Response) error {

	// Save our customer
	customer, err := h.repository.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Customer = customer
	return nil
}

// GetCustomers -
func (h *handler) GetCustomers(ctx context.Context, req *customerpb.Request, res *customerpb.Response) error {
	customers, _ := h.repository.GetAll()
	res.Customers = customers
	return nil
}

// getCustomerById -
func (h *handler) GetCustomerById(ctx context.Context, req *customerpb.Request, res *customerpb.Response) error {

	log.Println("Calling GetCustomerById")

	customer, _ := h.repository.GetById(req.AccountNumber)
	res.Customer = customer

	return nil
}

func (h *handler) GetCustomerAccounts(ctx context.Context, req *customerpb.Request, res *customerpb.Response) error {

	account := &accountpb.Account{
		Number: "NL57ABNA00000000",
		Type: "Savings",
		Status: "Active",
		Balance: 100,
	}


	_, err := h.accountClient.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

	//add code to call account service here *accountpb.Response
	accountResponse, err  := h.accountClient.GetAccountByNumber(ctx, &accountpb.Request{AccountNumber:req.AccountNumber})

	if err != nil {
		return err
	}

	if accountResponse == nil {
		log.Printf("No account found")
		return err
	}

	res.Account = accountResponse.Account
	return nil
}

