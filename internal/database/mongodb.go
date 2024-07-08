package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	client *mongo.Client
}

type MongoDbRepository struct {
	Id            string `bson:"_id,omitempty"`
	IsInitialized bool
}

func (mongodb *MongoDb) Connect() error {
	mongodbUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongodbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mongodb.client = client

	log.Println("connected to MongoDB")

	return nil
}

func (mongodb *MongoDb) Close() error {
	if err := mongodb.client.Disconnect(context.TODO()); err != nil {
		return err
	}

	log.Println("connection to MongoDB closed")

	return nil
}

func (mongodb *MongoDb) CreateRepository() (string, error) {
	repository := struct {
		IsInitialized bool
	}{
		IsInitialized: false,
	}

	insertResult, err := mongodb.getCollection().InsertOne(context.Background(), repository)

	if err != nil {
		return "", err
	}

	repositoryId := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return repositoryId, nil
}

func (mongodb *MongoDb) GetRepositoryById(id string) (*Repository, error) {
	var result MongoDbRepository

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objectId}}

	if err := mongodb.getCollection().FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &Repository{Id: result.Id, IsInitialized: result.IsInitialized}, nil
}

func (mongodb *MongoDb) DeleteRepository(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectId}

	_, err = mongodb.getCollection().DeleteOne(context.TODO(), filter)

	return err
}

func (mongodb *MongoDb) getCollection() *mongo.Collection {
	databaseName := os.Getenv("MONGODB_DATABASE_NAME")
	collectionName := os.Getenv("MONGODB_COLLECTION_NAME")

	return mongodb.client.Database(databaseName).Collection(collectionName)
}
