package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Alive!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", baseHandler)
	r.HandleFunc("/health", healthCheckHandler)
	http.ListenAndServe(":4000", r)
}
