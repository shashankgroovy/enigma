package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

var db_collection = os.Getenv("MONGO_COLLECTION")

// controller for rendering the homepage
func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// controller for health check
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Alive!")
}

// controller to create a secret message
func createSecretHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(guuid.New().String())
}

// controller to get a given secret message
func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secretHash := vars["hash"]
	json.NewEncoder(w).Encode(secretHash)
}

// controller to update a given secret message
func updateSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secretHash := vars["hash"]
	json.NewEncoder(w).Encode(secretHash)
}
