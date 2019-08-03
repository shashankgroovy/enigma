package server

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shashankgroovy/enigma/internal/utils"
)

// ConfigureRoutes setups up app routes and static routes
func ConfigureRoutes() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(utils.RequestLogger)

	// Health check
	r.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)

	// NOTE: It's important that subapps are served before the catch-all route ("/")
	// Setup sub apps or apis
	r = setupApiRoutes(r)

	// Serve static files
	r = setupStaticRoutes(r)

	// This should be the final route handling to catch-all requests and serve
	// our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(baseHandler("dist/index.html"))

	return r
}

// Creates routes for serving static assets
func setupStaticRoutes(r *mux.Router) *mux.Router {

	// Serve static files
	var dir string
	flag.StringVar(&dir, "dir", "dist", "The directory for static file content")
	flag.Parse()
	r.PathPrefix("/dist/").Handler(http.StripPrefix(
		"/dist/",
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
	// api.Use(mux.CORSMethodMiddleware(api))
	return r
}
