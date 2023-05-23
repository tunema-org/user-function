package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tunema-org/user-function/internal/backend"
	"github.com/tunema-org/user-function/internal/repository"
)

type RegisterInput struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	ProfileImgURL string `json:"profile_img_url"`
}

func (h *handler) Register(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var input RegisterInput

	if err := json.Unmarshal([]byte(req.Body), &input); err != nil {
		return JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
	}

	result, err := h.backend.Register(ctx, backend.RegisterParams{
		Username:      input.Username,
		Email:         input.Email,
		Password:      input.Password,
		ProfileImgURL: input.ProfileImgURL,
	})
	switch {
	case errors.Is(err, repository.ErrUserAlreadyExists):
		return JSON(http.StatusConflict, M{
			"message": "user already exists",
		})
	case err != nil:
		return JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
	}

	return JSON(http.StatusCreated, M{
		"message":      "user created",
		"access_token": result.AccessToken,
	})
}
