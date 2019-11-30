package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/harrietty/list-tweets/client"
	"github.com/harrietty/list-tweets/fetch"
	"github.com/harrietty/list-tweets/filter"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading environment variables", err)
	}

	c := client.New()

	tweets := fetch.Tweets(c, "harri_etty", "2019-01-01", "2019-01-01")
	filtered := filter.TweetsByKeyword(tweets, strings.ToLower("100DaysOfCode"))
	fmt.Printf("All tweets searched, found %v relevant tweets", strconv.Itoa(len(filtered)))
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
