package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type secret struct {
	Hash           string
	SecretText     string
	CreatedAt      int32 //time.Time
	ExpiresAt      int32 //time.Time
	RemainingViews int32
}

// Returns a single secret
func (s *secret) getSecret(col *mongo.Collection) error {
	return errors.New("Not implemented")
}

// Creates a new secret
func (s *secret) createSecret(col *mongo.Collection) error {
	return errors.New("Not implemented")
}

// Updates a secret
func (s *secret) updateSecret(col *mongo.Collection) error {
	return errors.New("Not implemented")
}

// Deletes a secret
func (s *secret) deleteSecret(col *mongo.Collection) error {
	return errors.New("Not implemented")
}
