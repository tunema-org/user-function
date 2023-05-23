package backend

import (
	"context"
	"time"

	"github.com/tunema-org/user-function/internal/jwt"
	"github.com/tunema-org/user-function/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterParams struct {
	Username      string
	Email         string
	Password      string
	ProfileImgURL string
}

type RegisterResult struct {
	UserID      int
	AccessToken string
}

func (b *Backend) Register(ctx context.Context, params RegisterParams) (RegisterResult, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResult{}, err
	}

	userID, err := b.repo.InsertUser(ctx, repository.User{
		Username:      params.Username,
		Email:         params.Email,
		Password:      string(hash),
		ProfileImgURL: params.ProfileImgURL,
	})
	if err != nil {
		return RegisterResult{}, err
	}

	accessToken, err := jwt.Generate(map[string]any{
		"userID": userID,
		"exp":    time.Now().Add(b.cfg.JWTDuration).Unix(),
		"email":  params.Email,
	}, b.cfg.JWTSecretKey)
	if err != nil {
		return RegisterResult{}, err
	}

	return RegisterResult{
		UserID:      userID,
		AccessToken: accessToken,
	}, nil
}
