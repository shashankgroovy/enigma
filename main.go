package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Setup the mux router
	r := mux.NewRouter()
	r.Use(requestLogger)

	// Index urls
	r.HandleFunc("/", baseHandler).Methods("GET")
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// Initialize the rest api
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/secret", createSecretHandler).Methods("POST")
	api.HandleFunc("/secret/{hash}", getSecretHandler).Methods("GET")

	// Serve static files
	var dir string
	flag.StringVar(&dir, "dir", "static", "The directory for static file content")
	flag.Parse()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// Get the port
	httpPort := os.Getenv("PORT")
	log.Printf("Server running on port %s\n", httpPort)

	// Setup server
	server := &http.Server{
		Handler:      r,
		Addr:         ":" + httpPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Fatal(server.ListenAndServe())
}
