package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tunema-org/user-function/internal/backend"
	"github.com/tunema-org/user-function/internal/repository"
)

type RegisterInput struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	ProfileImgURL string `json:"profile_img_url"`
}

func (h *handler) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	result, err := h.backend.Register(c.Request.Context(), backend.RegisterParams{
		Username:      input.Username,
		Email:         input.Email,
		Password:      input.Password,
		ProfileImgURL: input.ProfileImgURL,
	})
	switch {
	case errors.Is(err, repository.ErrUserAlreadyExists):
		c.JSON(http.StatusConflict, M{
			"message": "user already exists",
		})
		return
	case err != nil:
		log.Error().Err(err).Msg("problem with user register")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, M{
		"message":      "user created",
		"access_token": result.AccessToken,
	})
}
