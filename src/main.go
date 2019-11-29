package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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

	var finalID int64
	relevantTweets := []string{}
	shouldIncludeRetweets := false
	params := twitter.UserTimelineParams{
		IncludeRetweets: &shouldIncludeRetweets,
		ScreenName:      "harri_etty",
		Count:           200,
		TweetMode:       "extended",
	}

	foundAll := false

	// We can consider that all the tweets are found when the previous MaxID is the same as the current one
	for !foundAll {
		if finalID != 0 {
			params.MaxID = finalID
		}
		nextTweets := fetchTweetBatch(*client, params)
		newRelTweets := filterTweets(nextTweets)
		relevantTweets = append(relevantTweets, newRelTweets...)
		fmt.Printf("Found %v relevant tweets. Still searching...\n", len(newRelTweets))
		newFinalID := nextTweets[len(nextTweets)-1].ID
		if newFinalID == finalID {
			foundAll = true
		} else {
			finalID = newFinalID
		}
	}

	fmt.Println("All tweets searched, writing to file")

	writeToFile(relevantTweets)
}

func fetchTweetBatch(client twitter.Client, params twitter.UserTimelineParams) []twitter.Tweet {
	tweets, _, err := client.Timelines.UserTimeline(&params)
	if err != nil {
		fmt.Println(err)
	}
	return tweets
}

func filterTweets(tweets []twitter.Tweet) []string {
	res := []string{}
	for _, tweet := range tweets {
		if hasHashtag(tweet.Entities.Hashtags, "100DaysOfCode") {
			res = append(res, tweet.FullText)
		}
	}

	return res
}

func hasHashtag(tags []twitter.HashtagEntity, searchTerm string) bool {
	searchTerm = strings.ToLower(searchTerm)
	for _, t := range tags {
		if strings.ToLower(t.Text) == searchTerm {
			return true
		}
	}
	return false
}

func writeToFile(tweets []string) {
	data := []byte{}
	for _, line := range tweets {
		bts := []byte(line + string("\n------------------\n"))
		data = append(data, bts...)
	}
	err := ioutil.WriteFile("tweets.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
}
