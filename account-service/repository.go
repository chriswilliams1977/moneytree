package main

import (
	"context"

	pb "github.com/chriswilliams1977/account-service/proto/account"
	customerpb "github.com/chriswilliams1977/customer-service/proto/customer"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Create(customer *customerpb.Customer) error
	Get(customer And *customerpb.Customer) (account *pb.Account)
}

// VesselRepository ...
type AccountRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *AccountRepository) Get(customer *customerpb.Customer) (*pb.Account, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	var vessel *pb.Vessel
	if err := repository.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create a new vessel
func (repository *AccountRepository) Create(customer *customerpb.Customer) error {
	_, err := repository.collection.InsertOne(context.TODO(), vessel)
	return err
}
