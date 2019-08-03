package server

import (
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
	utils.RequestResponder(w, r, http.StatusOK, "Alive")
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

	utils.RequestResponder(w, r, http.StatusOK, "Alive")
}

// controller to get a given secret message
func getSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secretHash := vars["hash"]

	var sec models.Secret
	sec.Hash = secretHash

	err := sec.GetSecret()

	if err != nil {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	// Decrypt the secret message with secret hash
	secretText, err := utils.Decrypt([]byte(sec.SecretText), []byte(secretHash[:32]))

	if err != nil {
		log.Fatal("Couldn't decrypt text ", err)
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	sec.SecretText = string(secretText)
	if sec.RemainingViews < 0 {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}
		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}
	utils.RequestResponder(w, r, http.StatusOK, sec)
}

// controller to update a given secret message
func updateSecretHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	secretHash := vars["hash"]

	var sec models.Secret
	sec.Hash = secretHash

	err := sec.UpdateSecret()

	if err != nil {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	// Decrypt the secret message with secret hash
	secretText, err := utils.Decrypt([]byte(sec.SecretText), []byte(secretHash[:32]))

	if err != nil {
		log.Fatal("Couldn't decrypt text", err)
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}
		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	sec.SecretText = string(secretText)
	utils.RequestResponder(w, r, http.StatusOK, sec)
}

// controller to delete a given secret message
func deleteSecretHandler(w http.ResponseWriter, r *http.Request) {

	// Explicitly making it  since it's not required.
	log.Print("Delete operations are not permitted.")

	resp := utils.ErrorResponseObject{
		Status: http.StatusMethodNotAllowed,
		Error:  "Delete operations are not permitted",
	}
	utils.RequestResponder(w, r, http.StatusMethodNotAllowed, resp)
	return

	vars := mux.Vars(r)
	secretHash := vars["hash"]

	var sec models.Secret
	sec.Hash = secretHash

	sec.DeleteSecret()
	utils.RequestResponder(w, r, http.StatusOK, sec)
}
