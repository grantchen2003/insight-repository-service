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
	SessionId     string
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

func (mongodb *MongoDb) CreateRepository(sessionId string) (string, error) {
	repository := struct {
		SessionId     string
		IsInitialized bool
	}{
		SessionId:     sessionId,
		IsInitialized: false,
	}

	insertResult, err := mongodb.getCollection().InsertOne(context.Background(), repository)

	if err != nil {
		return "", err
	}

	repositoryId := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return repositoryId, nil
}

func (mongodb *MongoDb) GetRepositoryBySessionId(sessionId string) (*Repository, error) {
	var result MongoDbRepository

	filter := bson.D{{"sessionid", sessionId}}

	if err := mongodb.getCollection().FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &Repository{Id: result.Id, SessionId: result.SessionId, IsInitialized: result.IsInitialized}, nil
}

func (mongodb *MongoDb) SetRepositoryIsInitialized(id string, newIsInitialized bool) error {
	userDocId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", userDocId}}
	update := bson.D{
		{"$set", bson.D{
			{"isinitialized", newIsInitialized},
		}},
	}

	_, err = mongodb.getCollection().UpdateOne(context.TODO(), filter, update)
	return err
}

func (mongodb *MongoDb) getCollection() *mongo.Collection {
	databaseName := os.Getenv("MONGODB_DATABASE_NAME")
	collectionName := os.Getenv("MONGODB_COLLECTION_NAME")

	return mongodb.client.Database(databaseName).Collection(collectionName)
}
