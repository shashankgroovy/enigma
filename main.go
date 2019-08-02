package main

import (
	"context"
	"os"
	"time"
)

func main() {

	// Set a context with timeout for database interactions
	// Helps close db request when http requests go stale or are cancelled
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	// create an app
	app := App{}
	app.Initialize(ctx)
	app.Run(os.Getenv("PORT"))
}
