# List Tweets By Keyword

A quick script to fetch and filter tweets from your own timeline based on a search term.

#### Environment variables

This app uses Application-only authentication with OAuth2. You must have a Twitter Developer account to obtain application credentials.

The following environment variables must be set in a `.env` file:

    TWITTER_API_KEY=your_consumer_api_key
    TWITTER_SECRET_KEY=your_consumer_secret_api_key

Use the `.env.sample` as a template:

    cp .env.sample .env

#### Running the script

    go run src/main.go
