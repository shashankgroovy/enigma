package main

import (
	"context"
	"log"
	"net/http"
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

// Initializes a database and sets up routing
func (a *App) Initialize(ctx context.Context) {

	// Setup database
	client, err := models.ConfigureDB(ctx)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	testDbConnection(client)

	a.DB = client.Database(models.DB_NAME)

	// Setup routes
	a.Router = server.ConfigureRoutes()

}

// Starts a http.Server
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

// Pings the mongo database
func testDbConnection(client *mongo.Client) {

	// Check database connection
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Database connected!")
	}
}
