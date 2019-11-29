package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
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

	config := &clientcredentials.Config{
		ClientID:     consumerAPIKey,
		ClientSecret: consumerAPISecretKey,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(oauth2.NoContext)

	client := twitter.NewClient(httpClient)

	shouldIncludeRetweets := false

	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		IncludeRetweets: &shouldIncludeRetweets,
		ScreenName:      "harri_etty",
		Count:           500,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tweets[len(tweets)-1])
}
