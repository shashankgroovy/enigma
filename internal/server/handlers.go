package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/shashankgroovy/enigma/internal/models"
	"github.com/shashankgroovy/enigma/internal/utils"
)

// controller for rendering the homepage
func baseHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

// controller for health check
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Alive!")
}

// controller to create a secret message
func createSecretHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	secretText := r.FormValue("secretText")
	createdAt, _ := strconv.Atoi(r.FormValue("createdAt"))
	expiresAt, _ := strconv.Atoi(r.FormValue("expiresAt"))
	remainingViews, _ := strconv.Atoi(r.FormValue("remainingViews"))

	// Generate a unique uuid for hash
	uuid := guuid.New().String()

	// Encrypt the secret message with uuid as a 32 byte encryption key
	cipherText, err := utils.Encrypt([]byte(secretText), []byte(uuid[:32]))

	if err != nil {
		log.Fatal("Couldn't encrypt text", err)
	}

	sec := models.Secret{
		Hash:           uuid,
		SecretText:     string(cipherText),
		CreatedAt:      createdAt,
		ExpiresAt:      expiresAt,
		RemainingViews: remainingViews,
	}

	sec.CreateSecret()

	sec.SecretText = secretText
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sec)
}

// controller to get a given secret message
func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secretHash := vars["hash"]

	var sec models.Secret
	sec.Hash = secretHash

	sec.GetSecret()

	// Decrypt the secret message with secret hash
	secretText, err := utils.Decrypt([]byte(sec.SecretText), []byte(secretHash[:32]))

	if err != nil {
		log.Fatal("Couldn't decrypt text", err)
	}

	sec.SecretText = string(secretText)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sec)
}

// controller to update a given secret message
func updateSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secretHash := vars["hash"]

	var sec models.Secret
	sec.Hash = secretHash

	sec.UpdateSecret()

	// Decrypt the secret message with secret hash
	secretText, err := utils.Decrypt([]byte(sec.SecretText), []byte(secretHash[:32]))

	if err != nil {
		log.Fatal("Couldn't decrypt text", err)
	}

	sec.SecretText = string(secretText)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sec)
}

// controller to delete a given secret message
func deleteSecretHandler(w http.ResponseWriter, r *http.Request) {
	// Explicitly making it  since it's not required.
	log.Print("Delete operations are not permitted.")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w)
	return

	vars := mux.Vars(r)
	secretHash := vars["hash"]

	var sec models.Secret
	sec.Hash = secretHash

	sec.DeleteSecret()
	json.NewEncoder(w).Encode(sec)
}
