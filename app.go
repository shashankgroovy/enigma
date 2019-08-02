package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"enigma/internal/models"
	"enigma/internal/server"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func (a *App) Initialize(ctx context.Context) {

	client, err := models.ConfigureDB(ctx)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	testDbConnection(client)

	a.DB = client.Database(os.Getenv("MONGO_DBNAME"))

	// Setup routes
	a.Router = server.ConfigureRoutes()

}

func (a *App) Run(httpPort string) {

	log.Printf("Server running on port %s\n", httpPort)

	// Setup server
	server := &http.Server{
		Handler:      a.Router,
		Addr:         ":" + httpPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

// pings the mongo database
func testDbConnection(client *mongo.Client) {

	// Check database connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}
