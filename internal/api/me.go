package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *handler) Me(c *gin.Context) {
	authorization := strings.Split(c.Request.Header["Authorization"][0], " ")
	if len(authorization) != 2 {
		c.JSON(http.StatusUnauthorized, M{
			"message": "please login",
		})
		return
	}

	result, err := h.backend.Me(c.Request.Context(), authorization[1])
	if err != nil {
		log.Error().Err(err).Msg("problem with user me")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
