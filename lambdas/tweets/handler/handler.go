package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/harrietty/list-tweets/fetch"
	"github.com/harrietty/list-tweets/filter"
	"regexp"
	"time"
)

// Handler struct
type Handler struct {
	stage  string
	client *twitter.Client
}

// New creates a new handler
func New(stage string, client *twitter.Client) Handler {
	return Handler{
		stage:  stage,
		client: client,
	}
}

// HandleRequest handles incoming API gateway requests
func (h Handler) HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	username := request.QueryStringParameters["username"]
	dateSinceStr := request.QueryStringParameters["date_since"]
	dateBeforeStr := request.QueryStringParameters["date_before"]
	filterString := request.QueryStringParameters["filter"]
	if username == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "username is required"}, nil
	}

	if filterString == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "filter is required"}, nil
	}

	dateSince, dateBefore := "", ""
	if dateSinceStr != "" {
		ds, err := time.Parse("2006-01-02", dateSinceStr)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Cannot parse dateSince"}, nil
		}
		dateSince = ds.Format("2006-01-02")
	}
	if dateBeforeStr != "" {
		db, err := time.Parse("2006-01-02", dateBeforeStr)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Cannot parse dateBefore"}, nil
		}
		dateBefore = db.Format("2006-01-02")
	}

	// Make a twitter API call
	tw, err := fetch.Tweets(h.client, username, dateSince, dateBefore)

	// Handle Username not found/other Twitter API errors
	if err != nil {
		matched, _ := regexp.MatchString("34", err.Error())
		if matched {
			return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Username not found"}, nil
		}
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Could not fetch Tweets"}, nil
	}

	// Filter the tweets by filterString
	filtered := filter.TweetsByKeyword(tw, filterString)
	// Respond with the slice of tweets
	blob, err := json.Marshal(filtered)
	if err != nil {
		fmt.Println("Cannot parse filtered tweets as JSON", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error parsing filtered tweets as JSON"}, nil
	}

	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Credentials"] = "true"
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(blob), Headers: headers}, nil
}
