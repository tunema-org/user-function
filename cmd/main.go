package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tunema-org/user-function/internal/api"
	"github.com/tunema-org/user-function/internal/backend"
	"github.com/tunema-org/user-function/internal/clients"
	"github.com/tunema-org/user-function/internal/config"
	"github.com/tunema-org/user-function/internal/repository"
)

func main() {
	// region := os.Getenv("AWS_REGION")
	// awsSession, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(region),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	clients, err := clients.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(clients.DB)
	backend := backend.New(clients, repo, cfg)
	handler := api.NewHandler(ctx, backend)

	lambda.Start(handler)
}
