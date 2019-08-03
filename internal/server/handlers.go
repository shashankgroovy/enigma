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
	utils.RequestResponder(w, r, http.StatusOK, "Alive")
}

// controller to create a secret message
func createSecretHandler(w http.ResponseWriter, r *http.Request) {

	// declare the Secret
	var sec models.Secret
	var secretText string

	switch r.Header.Get("Content-Type") {

	case "application/json":
		// Content-Type: application/json
		// Populate the struct from JSON Data
		err := json.NewDecoder(r.Body).Decode(&sec)
		if err != nil {
			log.Printf("decoded %+v", sec)
		}
		log.Printf("decoded 2 %+v", sec)

	default:
		// Content-Type: application/x-www-form-urlencoded
		// Populate the struct with Form Data

		// Parse form data
		r.ParseForm()
		secretText = r.FormValue("secretText")
		createdAt, _ := strconv.Atoi(r.FormValue("createdAt"))
		expiresAt, _ := strconv.Atoi(r.FormValue("expiresAt"))
		remainingViews, _ := strconv.Atoi(r.FormValue("remainingViews"))

		sec.SecretText = secretText
		sec.CreatedAt = createdAt
		sec.ExpiresAt = expiresAt
		sec.RemainingViews = remainingViews
	}

	if sec.SecretText == "" {
		resp := utils.ErrorResponseObject{
			Status: http.StatusBadRequest,
			Error:  "Not enough parameters",
		}
		log.Print("Not enough parameters")
		utils.RequestResponder(w, r, http.StatusBadRequest, resp)
		return
	}

	// Generate a unique uuid for hash
	uuid := guuid.New().String()

	// Encrypt the secret message with uuid as a 32 byte encryption key
	cipherText, err := utils.Encrypt([]byte(secretText), []byte(uuid[:32]))

	if err != nil {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}
		log.Print("Couldn't encrypt text", err)

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	// Populate secret struct
	sec.Hash = uuid
	sec.SecretText = string(cipherText)

	sec.CreateSecret()

	utils.RequestResponder(w, r, http.StatusOK, sec)
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
		log.Print("Couldn't decrypt text ", err)
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	sec.SecretText = string(secretText)
	if sec.RemainingViews <= 0 {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  "Expired secret!",
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

	err := sec.GetSecret()
	if err != nil {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	} else if sec.RemainingViews < 0 {
		resp := utils.ErrorResponseObject{
			Status: http.StatusNotFound,
			Error:  "Expired secret!",
		}

		utils.RequestResponder(w, r, http.StatusNotFound, resp)
		return
	}

	err = sec.UpdateSecret()

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
		log.Print("Couldn't decrypt text", err)
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
