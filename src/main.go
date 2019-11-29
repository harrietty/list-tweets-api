package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading environment variables", err)
	}

	consumerAPIKey, exists := os.LookupEnv("TWITTER_API_KEY")
	if !exists {
		fmt.Println("TWITTER_API_KEY must be set")
	}
	consumerAPISecretKey, exists := os.LookupEnv("TWITTER_SECRET_KEY")
	if !exists {
		fmt.Println("TWITTER_SECRET_KEY must be set")
	}

	fmt.Println(consumerAPIKey, consumerAPISecretKey)
}
