package fetch

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

// Tweets fetches all a user's tweets between two given date strings
func Tweets(client *twitter.Client, username string, dateSince string, dateBefore string) []twitter.Tweet {
	var finalID int64
	tw := []twitter.Tweet{}
	shouldIncludeRetweets := false
	params := twitter.UserTimelineParams{
		IncludeRetweets: &shouldIncludeRetweets,
		ScreenName:      username,
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
		fmt.Printf("Found %v tweets. Still searching...\n", len(tw))
		tw = append(tw, nextTweets...)
		newFinalID := nextTweets[len(nextTweets)-1].ID
		if newFinalID == finalID {
			foundAll = true
		} else {
			finalID = newFinalID
		}
	}

	return tw
}

func fetchTweetBatch(client twitter.Client, params twitter.UserTimelineParams) []twitter.Tweet {
	tweets, _, err := client.Timelines.UserTimeline(&params)
	if err != nil {
		fmt.Println(err)
	}
	return tweets
}
