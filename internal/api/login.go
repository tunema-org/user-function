package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tunema-org/user-function/internal/backend"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *handler) Login(c *gin.Context) {
	var input LoginInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
	}

	result, err := h.backend.Login(c.Request.Context(), backend.LoginParams{
		Email:    input.Email,
		Password: input.Password,
	})
	switch {
	case errors.Is(err, backend.ErrLoginInvalidCredentials):
		c.JSON(http.StatusUnauthorized, M{
			"message": backend.ErrLoginInvalidCredentials.Error(),
		})
	case err != nil:
		log.Error().Err(err).Msg("problem with user login")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
	default:
		c.JSON(http.StatusCreated, M{
			"message":      "user logged in",
			"access_token": result.AccessToken,
		})
	}
}
