package client

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// New initialises and returns a new Twitter client with Application-only authentication
func New() *twitter.Client {
	consumerAPIKey, exists := os.LookupEnv("TWITTER_API_KEY")
	if !exists {
		fmt.Println("TWITTER_API_KEY must be set")
	}
	consumerAPISecretKey, exists := os.LookupEnv("TWITTER_SECRET_KEY")
	if !exists {
		fmt.Println("TWITTER_SECRET_KEY must be set")
	}

	config := &clientcredentials.Config{
		ClientID:     consumerAPIKey,
		ClientSecret: consumerAPISecretKey,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)

	return twitter.NewClient(httpClient)
}
