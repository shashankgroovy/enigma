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
	DB            *mongo.Database
	DB_HOST       string = os.Getenv("MONGO_HOST")
	DB_NAME       string = os.Getenv("MONGO_DBNAME")
	DB_PASS       string = os.Getenv("MONGO_PASSWORD")
	DB_USER       string = os.Getenv("MONGO_USER")
	DB_COLLECTION string = os.Getenv("MONGO_COLLECTION")
)

// ConfigureDB initializes a mongo connection client and returns it
func ConfigureDB(ctx context.Context) (*mongo.Client, error) {
	connectionString := fmt.Sprintf(`mongodb://%s:%s@%s/%s`,
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_NAME,
	)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}

	DB = client.Database(DB_NAME)
	return client, nil

}

// TestDbConnection tests mongo connection by sending pings to the mongo
// database
func TestDbConnection(client *mongo.Client) {

	// Check database connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}

// GetDefaultCollection returns the "secret" collection
func GetDefaultCollection(db *mongo.Database) (col *mongo.Collection) {
	col = db.Collection(DB_COLLECTION)
	return
}
