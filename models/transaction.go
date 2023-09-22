package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Reference       string             `bson:"reference"`
	Description     string             `json:"description"`
	TransactionType string             `json:"transaction_type"`
	Amount          float64            `json:"amount"`
}
