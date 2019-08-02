package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shashankgroovy/enigma/internal/models"
	"github.com/shashankgroovy/enigma/internal/server"
	"go.mongodb.org/mongo-driver/mongo"
)

// App struct for keeping things simple and concise.
type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

// Initialize method initializes a database and sets up routing
func (a *App) Initialize(ctx context.Context) {

	// Setup database
	client, err := models.ConfigureDB(ctx)

	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	models.TestDbConnection(client)

	a.DB = client.Database(models.DB_NAME)

	// Setup routes
	a.Router = server.ConfigureRoutes()

}

// Run method starts a http.Server
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
