package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string, databaseName string, collectioName string) *mongo.Database {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}
	db := client.Database(databaseName)

	return db

}

func Disconnect(client *mongo.Client) {
	client.Disconnect(context.TODO())
}
