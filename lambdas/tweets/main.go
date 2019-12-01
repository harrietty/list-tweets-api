package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/harrietty/list-tweets/client"
	"github.com/harrietty/list-tweets/lambdas/tweets/handler"
	"os"
)

func main() {
	stage, exists := os.LookupEnv("STAGE")
	if !exists {
		stage = "dev"
	}

	c := client.New()
	h := handler.New(stage, c)
	lambda.Start(h.HandleRequest)
}
