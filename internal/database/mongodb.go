package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func (mongodb *MongoDB) Connect() error {
	mongodbUri := os.Getenv("MONGODB_URI")

	// Connect to the database.
	clientOptions := options.Client().ApplyURI(mongodbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	mongodb.client = client

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB")

	return err
}

func (mongodb *MongoDB) Close() error {
	err := mongodb.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	fmt.Println("Connection closed.")

	return err
}

func (mongodb *MongoDB) BatchSave(databaseName string, collectionName string, documentBatch []interface{}) error {
	collection := mongodb.client.Database(databaseName).Collection(collectionName)

	_, err := collection.InsertMany(context.Background(), documentBatch)

	return err
}
