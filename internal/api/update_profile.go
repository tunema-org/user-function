package api

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tunema-org/user-function/internal/backend"
	"github.com/tunema-org/user-function/internal/mime"
)

type UpdateProfileInput struct {
	Username string `form:"username" binding:"required"`
}

func (h *handler) UpdateProfile(c *gin.Context) {
	var input UpdateProfileInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	profileImgFileHeader, err := c.FormFile("profile_img")
	if err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	profileImg, err := profileImgFileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid request body",
		})
		return
	}

	// max file size 5mb
	if profileImgFileHeader.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, M{
			"message": "max profile image size is 5mb",
		})
		return
	}

	acceptedImageFileTypes := []string{"image/jpeg", "image/png"}

	if !mime.Contains(profileImg, acceptedImageFileTypes) {
		c.JSON(http.StatusBadRequest, M{
			"message": "invalid profile image type",
		})
		return
	}

	authorization := strings.Split(c.Request.Header["Authorization"][0], " ")
	if len(authorization) != 2 {
		c.JSON(http.StatusUnauthorized, M{
			"message": "please login",
		})
		return
	}

	err = h.backend.UpdateProfile(c.Request.Context(), authorization[1], backend.UpdateProfileParams{
		Username:       input.Username,
		ProfileImg:     profileImg,
		ProfileImgType: filepath.Ext(profileImgFileHeader.Filename),
	})
	if err != nil {
		log.Err(err).Msg("problem with user update profile")
		c.JSON(http.StatusInternalServerError, M{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, M{
		"message": "profile updated",
	})
}
