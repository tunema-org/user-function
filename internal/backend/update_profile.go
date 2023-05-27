package backend

import (
	"context"
	"errors"
	"io"
	"path"

	"github.com/google/uuid"
	"github.com/tunema-org/user-function/internal/jwt"
	"github.com/tunema-org/user-function/internal/repository"
)

type UpdateProfileParams struct {
	Username       string
	ProfileImg     io.Reader
	ProfileImgType string
}

func (b *Backend) UpdateProfile(ctx context.Context, accessToken string, params UpdateProfileParams) error {
	_, claims, err := jwt.Verify(accessToken, b.cfg.JWTSecretKey)
	if err != nil {
		return err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return errors.New("invalid claims")
	}

	key := path.Join("users", uuid.New().String()+params.ProfileImgType)

	output, err := b.clients.S3.UploadFile(ctx, key, params.ProfileImg)
	if err != nil {
		return err
	}

	err = b.repo.UpdateUser(ctx, int(userID), repository.User{
		Username:      params.Username,
		ProfileImgURL: output.Location,
	})
	if err != nil {
		return err
	}

	return nil
}
