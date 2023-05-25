package backend

import (
	"context"
	"errors"

	"github.com/tunema-org/user-function/internal/jwt"
)

type MeResult struct {
	UserID        int    `json:"user_id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	ProfileImgURL string `json:"profile_img_url"`
}

func (b *Backend) Me(ctx context.Context, accessToken string) (MeResult, error) {
	_, claims, err := jwt.Verify(accessToken, b.cfg.JWTSecretKey)
	if err != nil {
		return MeResult{}, err
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return MeResult{}, errors.New("invalid claims")
	}

	user, err := b.repo.FindUserByID(ctx, int(userID))
	if err != nil {
		return MeResult{}, err
	}

	return MeResult{
		UserID:        user.ID,
		Username:      user.Username,
		Email:         user.Email,
		ProfileImgURL: user.ProfileImgURL,
	}, nil
}
