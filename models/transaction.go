package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" form:"id"`
	Reference       string             `bson:"reference" form:"ref"`
	Description     string             `json:"description" form:"desc"`
	TransactionType TransactionType    `json:"transaction_type" form:"trantype"`
	Amount          float64            `json:"amount" form:"amount"`
}
