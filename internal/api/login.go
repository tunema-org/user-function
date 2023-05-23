package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tunema-org/user-function/internal/backend"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) Login(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var input LoginInput

	if err := json.Unmarshal([]byte(req.Body), &input); err != nil {
		return JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
	}

	result, err := h.backend.Login(ctx, backend.LoginParams{
		Email:    input.Email,
		Password: input.Password,
	})
	switch {
	case errors.Is(err, backend.ErrLoginInvalidCredentials):
		return JSON(http.StatusUnauthorized, M{
			"message": backend.ErrLoginInvalidCredentials.Error(),
		})
	case err != nil:
		return JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
	}

	return JSON(http.StatusCreated, M{
		"message":      "user logged in",
		"access_token": result.AccessToken,
	})
}
