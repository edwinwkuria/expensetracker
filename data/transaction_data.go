package data

import (
	"context"
	"expensetracker/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	// Access the database and collection
	db := client.Database("expensetracker")
	return db.Collection("transaction")
}

func SaveTransaction(m models.Transaction) bool {

	collection := GetCollection()
	//query db
	result, error := collection.InsertOne(context.TODO(), m)
	if error != nil {
		return false
	}
	fmt.Printf("Returns %v\n", result.InsertedID)
	return true

}

func GetTransactions() []models.Transaction {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	collection := GetCollection()
	//query db
	cursor, error := collection.Find(ctx, bson.D{})
	if error != nil {
		fmt.Println(error)
	}

	defer cursor.Close(ctx)

	// Loop through the documents and process them
	var results []models.Transaction
	if err := cursor.All(ctx, &results); err != nil {
		fmt.Println(err)
	}
	// return
	return results

}

func GetTransaction(ID string) models.Transaction {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	prim, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		fmt.Println("Error converting ID:", err)
		fmt.Println(err)
	}
	collection := GetCollection()
	filter := bson.D{{"_id", prim}}
	//query db
	var results models.Transaction
	error := collection.FindOne(ctx, filter).Decode(&results)
	if error != nil {
		fmt.Println(error)
	}
	// return
	return results

}

func UpdateTransaction(id string, m models.Transaction) bool {

	collection := GetCollection()
	prim, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error converting ID:", err)
		fmt.Println(err)
	}
	filter := bson.D{{"_id", prim}}
	update := bson.M{"$set": m}
	//query db
	result, error := collection.UpdateOne(context.TODO(), filter, update)
	if error != nil {
		return false
	}
	fmt.Printf("Returns %v\n", result)
	return true

}

func DeleteTransaction(m string) bool {

	collection := GetCollection()
	prim, err := primitive.ObjectIDFromHex(m)
	if err != nil {
		fmt.Println("Error converting ID:", prim)
		fmt.Println(err)
	}
	filter := bson.D{{"_id", prim}}
	//query db
	result, error := collection.DeleteOne(context.TODO(), filter)
	if error != nil {
		return false
	}
	fmt.Printf("Returns %v\n", result)
	return true

}
