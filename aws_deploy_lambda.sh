#!/bin/bash

set -ex

# Get the current date in the desired format
now=$(date +"%m-%d-%Y-%H-%M-%S")

# Get the current Git commit hash
commit_hash=$(git rev-parse --short HEAD)

# Create the zip file name by combining the date and commit hash
build_name="${now}_deployment_${commit_hash}"

# Build golang executable
echo "[1] Building golang executable..."
GOOS=linux GOARCH=amd64 go build -o build/main ./cmd/main.go

# Zip golang executable
echo "[2] Zipping golang executable..."

zip -jr build/$build_name.zip build/main

# Deploy to AWS Lambda
echo "[3] Deploying to AWS Lambda..."
aws lambda update-function-code --function-name tunema-userFunction --zip-file fileb://build/$build_name.zip
