package filter

import (
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

// TweetsByKeyword takes an array of tweets and filters by a given keyword
func TweetsByKeyword(tweets []twitter.Tweet, searchTerm string) []twitter.Tweet {
	res := []twitter.Tweet{}
	for _, tweet := range tweets {
		if hasKeyword(tweet.Entities.Hashtags, searchTerm) {
			res = append(res, tweet) // Could return Tweet.FullText
		}
	}
	return res
}

func hasKeyword(tags []twitter.HashtagEntity, searchTerm string) bool {
	for _, t := range tags {
		if strings.ToLower(t.Text) == searchTerm {
			return true
		}
	}
	return false
}
