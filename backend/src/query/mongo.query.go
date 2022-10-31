package query

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to mongodb
func ConnectMongo(mongoUsername string, mongoPassword string) *mongo.Client {
	// Authentication
	credentials := options.Credential{
		AuthSource: "admin",
		Username:   mongoUsername,
		Password:   mongoPassword,
	}
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").
		SetAuth(credentials)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
