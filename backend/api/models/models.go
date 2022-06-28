package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Payment struct {
		ID     primitive.ObjectID `bson:"id"`
		ItemID string             `json:"item_id"`
		Paid   bool               `json:"paid"`
	}

	Items struct {
		ID          primitive.ObjectID `bson:"id"`
		Name        string             `bson:"item_name"`
		Price       float64            `bson:"price"`
		Discount    float64            `bson:"discount"`
		Description string             `bson:"description"`
		Tags        []string           `bson:"tag"`
		Image       string             `bson:"image"`
	}

	Delivery struct {
		ID        primitive.ObjectID `bson:"id"`
		FirstName string             `bson:"first_name"`
		LastName  string             `bson:"last_name"`
		Address   string             `bson:"address"`
	}
)
