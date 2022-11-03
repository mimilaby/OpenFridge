// Package query handles the queries to the api
package query

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Food struct holds the food data
type Food struct {
	Name string `json:"name"`
}

// GetFood get food from mongoDB
func GetFood(client *mongo.Client, name string) (Food, error) {
	var food Food
	collection := client.Database("openfridge").Collection("food")
	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&food)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(food)
	return food, nil
}
