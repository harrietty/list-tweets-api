#!/usr/bin/env bash

sudo apt-get update
sudo apt-get install -y nodejs npm
npm install -g serverless
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh