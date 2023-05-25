package api

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

func (h *handler) Me(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	accessToken, ok := req.Headers["Authorization"]
	if !ok {
		return JSON(http.StatusUnauthorized, M{
			"message": "please login",
		})
	}

	result, err := h.backend.Me(ctx, accessToken)
	if err != nil {
		log.Error().Err(err).Msg("problem with user me")
		return JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
	}

	return JSON(http.StatusOK, result)
}
