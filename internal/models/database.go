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
	DB_HOST       = os.Getenv("MONGO_HOST")
	DB_NAME       = os.Getenv("MONGO_DBNAME")
	DB_PASS       = os.Getenv("MONGO_PASSWORD")
	DB_USER       = os.Getenv("MONGO_USER")
	DB_COLLECTION = os.Getenv("MONGO_COLLECTION")
)

// configure and setup mongo
func ConfigureDB(ctx context.Context) (*mongo.Client, error) {
	connection_string := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_NAME,
	)

	clientOptions := options.Client().ApplyURI(connection_string)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	return client, nil

}

// returns the secret collection
func GetDefaultCollection(db *mongo.Database) (col *mongo.Collection) {
	col = db.Collection(DB_COLLECTION)
	return
}
