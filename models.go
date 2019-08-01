package main

import "time"

type Secret struct {
	Hash           string
	SecretText     string
	CreatedAt      time.Time
	ExpiresAt      time.Time
	RemainingViews int32
}

var secrets []Secret
