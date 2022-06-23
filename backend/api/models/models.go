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
		Name        string             `json:"item_name"`
		Price       float64            `json:"price"`
		Discount    float64            `json:"discount"`
		Description string             `json:"description"`
		Tag         []string           `json:"tag"`
		Image       string             `json:"image"`
	}

	Delivery struct {
		ID        primitive.ObjectID `bson:"id"`
		FirstName string             `json:"first_name"`
		LastName  string             `json:"last_name"`
		Address   string             `json:"address"`
	}
)
