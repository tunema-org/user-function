package backend

import (
	"context"
	"time"

	"github.com/tunema-org/user-function/internal/jwt"
	"github.com/tunema-org/user-function/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterParams struct {
	Username      string
	Email         string
	Password      string
	ProfileImgURL string
}

type RegisterResult struct {
	User        model.User
	AccessToken string
}

func (b *Backend) Register(ctx context.Context, params RegisterParams) (RegisterResult, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResult{}, err
	}

	user, err := b.repo.InsertUser(ctx, model.User{
		Username:      params.Username,
		Email:         params.Email,
		Password:      string(hash),
		ProfileImgURL: params.ProfileImgURL,
	})
	if err != nil {
		return RegisterResult{}, err
	}

	accessToken, err := jwt.Generate(map[string]any{
		"userID": user.ID,
		"exp":    time.Now().Add(b.cfg.JWTDuration).Unix(),
		"email":  params.Email,
	}, b.cfg.JWTSecretKey)
	if err != nil {
		return RegisterResult{}, err
	}

	return RegisterResult{
		User:        user,
		AccessToken: accessToken,
	}, nil
}
