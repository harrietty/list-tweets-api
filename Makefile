build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/tweets lambdas/tweets/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy:
	sls deploy --verbose --TWITTER_API_KEY=${TWITTER_API_KEY} --TWITTER_SECRET_KEY=${TWITTER_SECRET_KEY}

deployprod:
	sls deploy --verbose --stage prod