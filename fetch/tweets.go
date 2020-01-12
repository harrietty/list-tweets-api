package fetch

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

// Tweets fetches all a user's tweets between two given date strings
func Tweets(client *twitter.Client, username string, dateSince string, dateBefore string) ([]twitter.Tweet, error) {
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
		nextTweets, err := fetchTweetBatch(*client, params)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Found %v tweets. Still searching...\n", len(tw))
		tw = append(tw, nextTweets...)

		// If the user has no tweets, no point in continuing
		if len(tw) == 0 {
			return tw, nil
		}

		newFinalID := nextTweets[len(nextTweets)-1].ID
		if newFinalID == finalID {
			foundAll = true
		} else {
			finalID = newFinalID
		}
	}

	return tw, nil
}

func fetchTweetBatch(client twitter.Client, params twitter.UserTimelineParams) ([]twitter.Tweet, error) {
	tweets, _, err := client.Timelines.UserTimeline(&params)
	if err != nil {
		fmt.Println("Error fetching tweets", err)
		return tweets, err
	}
	return tweets, nil
}
