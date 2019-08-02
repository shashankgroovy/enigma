package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Struct for containing secret
type Secret struct {
	Hash           string `json:"hash" bson:"hash"`
	SecretText     string `json:"secretText" bson:"secretText"`
	CreatedAt      int    `json:"createdAt" bson:"createsAt"`
	ExpiresAt      int    `json:"expiresAt" bson:"expiresAt"`
	RemainingViews int    `json:"remainingViews" bson:"remainingViews"`
}

// CreateSecret method creates a new secret
func (s *Secret) CreateSecret() {
	collection := GetDefaultCollection(DB)

	// Execute the query
	_, err := collection.InsertOne(context.Background(), s)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Created a secret! ", s.Hash)
}

// GetSecret method fetches a secret
func (s *Secret) GetSecret() {
	collection := GetDefaultCollection(DB)

	// conditions
	filter := bson.D{{"hash", s.Hash}}

	// Execute the query
	err := collection.FindOne(context.Background(), filter).Decode(&s)

	if err != nil {
		log.Print("Secret not found.", err)
	}
	log.Print("Found a secret! ", s.Hash)
}

// UpdateSecret method updates a given secret
func (s *Secret) UpdateSecret() {
	collection := GetDefaultCollection(DB)

	// conditions
	filter := bson.D{{"hash", s.Hash}}
	update := bson.D{
		{"$inc", bson.D{
			{"remainingViews", -1},
		}},
	}
	// Execute the query
	_, err := collection.UpdateOne(context.Background(), filter, update)
	err = collection.FindOne(context.Background(), filter).Decode(&s)

	if err != nil {
		log.Print("Secret not found.", err)
	}
	log.Print("Updated a secret! ", s.Hash)
}

// DeleteSecrets method deletes a given secret
func (s *Secret) DeleteSecret() {
	collection := GetDefaultCollection(DB)

	// conditions
	filter := bson.D{{"hash", s.Hash}}

	// Execute the query
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Print("Secret not found.", err)
	}
	log.Print("Deleted a secret!")
}
