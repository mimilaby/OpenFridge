package query

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Food struct {
	Name string `json:"name"`
}

func GetFood(client *mongo.Client, name string) (Food, error) {
	collection := client.Database("food").Collection("food")
	filter := bson.D{{"name", name}}
	var result Food
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found a single document: ", result)
	return result, nil
}
