package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/tunema-org/user-function/internal/api"
	"github.com/tunema-org/user-function/internal/backend"
	"github.com/tunema-org/user-function/internal/clients"
	"github.com/tunema-org/user-function/internal/config"
	"github.com/tunema-org/user-function/internal/repository"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New()

	clients, err := clients.New(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	repo := repository.New(clients.DB)
	backend := backend.New(clients, repo, cfg)
	handler := api.NewHandler(ctx, backend)

	if !isLambda() {
		handler.Run(cfg.Address)
		return
	}

	log.Info().Msg("Running in AWS Lambda")
	lambdaHandler := createLambdaHandler(ctx, handler)
	lambda.Start(lambdaHandler)
}

type LambdaHandler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func createLambdaHandler(ctx context.Context, api *gin.Engine) LambdaHandler {
	ginLambda := ginadapter.New(api)
	return func(ctx_ context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return ginLambda.ProxyWithContext(ctx_, req)
	}
}

func isLambda() bool {
	return os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != ""
}
