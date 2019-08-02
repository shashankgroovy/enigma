package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db_host       = os.Getenv("MONGO_HOST")
	db_name       = os.Getenv("MONGO_DBNAME")
	db_pass       = os.Getenv("MONGO_PASSWORD")
	db_user       = os.Getenv("MONGO_USER")
	db_collection = os.Getenv("MONGO_COLLECTION")
)

// configure and setup mongo
func ConfigureDB(ctx context.Context) (*mongo.Client, error) {
	db_uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		db_user,
		db_pass,
		db_host,
		db_name,
	)

	clientOptions := options.Client().ApplyURI(db_uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	return client, nil

}

// returns the secret collection
func GetDefaultCollection(db *mongo.Database) (col *mongo.Collection) {
	col = db.Collection(db_collection)
	return
}
