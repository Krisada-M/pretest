package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB for connect mongo database
func ConnectDB() *mongo.Client {
	EnvLoad()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(os.Getenv("DB_URI")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("Employee").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	fmt.Println("----- Connect successfully to MongoDB! -----")
	return client
}

// SelectCollection for select collection in DB
func SelectCollection(client *mongo.Client, collectionname string) *mongo.Collection {
	collection := client.Database(os.Getenv("DB_NAME")).Collection(collectionname)
	return collection
}
