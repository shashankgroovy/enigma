package server

import (
	"flag"
	"net/http"

	"enigma/internal/utils"

	"github.com/gorilla/mux"
)

// Setup application routes
func ConfigureRoutes() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(utils.RequestLogger)

	// Index urls
	r.HandleFunc("/", baseHandler).Methods("GET")
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	r = createStaticRoutes(r)
	r = createApiRoutes(r)
	return r
}

// Creates routes for serving static assets
func createStaticRoutes(r *mux.Router) *mux.Router {

	// Serve static files
	var dir string
	flag.StringVar(&dir, "dir", "static", "The directory for static file content")
	flag.Parse()
	r.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir(dir))),
	)

	return r

}

// Creates routes for rest api
func createApiRoutes(r *mux.Router) *mux.Router {

	// Initialize the rest api
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/secret", createSecretHandler).Methods("POST")
	api.HandleFunc("/secret/{hash}", getSecretHandler).Methods("GET")
	api.HandleFunc("/secret/{hash}", updateSecretHandler).Methods("PUT")
	return r
}
