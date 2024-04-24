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

func (mongodb *MongoDB) Save(databaseName string, collectionName string, document interface{}) error {
	collection := mongodb.client.Database(databaseName).Collection(collectionName)

	insertResult, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		return err
	}

	fmt.Println("Inserted document with ID:", insertResult.InsertedID)

	return err
}

func (mongodb *MongoDB) BatchSave(databaseName string, collectionName string, documentBatch []interface{}) error {
	collection := mongodb.client.Database(databaseName).Collection(collectionName)

	insertResults, err := collection.InsertMany(context.Background(), documentBatch)

	if err != nil {
		return err
	}

	for _, id := range insertResults.InsertedIDs {
		fmt.Println("Inserted document with ID:", id)
	}

	return err
}
