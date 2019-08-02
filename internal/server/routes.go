package server

import (
	"flag"
	"net/http"

	"enigma/internal/utils"

	"github.com/gorilla/mux"
)

// ConfigureRoutes setups up app routes and static routes
func ConfigureRoutes() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(utils.RequestLogger)

	// Index urls
	r.HandleFunc("/", baseHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)

	r = setupStaticRoutes(r)
	r = setupApiRoutes(r)

	return r
}

// Creates routes for serving static assets
func setupStaticRoutes(r *mux.Router) *mux.Router {

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
func setupApiRoutes(r *mux.Router) *mux.Router {

	// Initialize the rest api
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/secret", createSecretHandler).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/secret/{hash}", getSecretHandler).Methods(http.MethodGet)
	api.HandleFunc("/secret/{hash}", updateSecretHandler).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/secret/{hash}", deleteSecretHandler).Methods(http.MethodDelete)

	// CORS
	api.Use(mux.CORSMethodMiddleware(api))
	return r
}
