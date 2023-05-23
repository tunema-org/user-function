package backend

import (
	"context"
	"errors"
	"time"

	"github.com/tunema-org/user-function/internal/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginParams struct {
	Email    string
	Password string
}

type LoginResult struct {
	UserID      int
	AccessToken string
}

var (
	ErrLoginInvalidCredentials = errors.New("invalid credentials")
)

func (b *Backend) Login(ctx context.Context, params LoginParams) (LoginResult, error) {
	user, err := b.repo.FindUserByEmail(ctx, params.Email)
	if err != nil {
		return LoginResult{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return LoginResult{}, ErrLoginInvalidCredentials
	}

	accessToken, err := jwt.Generate(map[string]any{
		"userID": user.ID,
		"exp":    time.Now().Add(b.cfg.JWTDuration).Unix(),
		"email":  params.Email,
	}, b.cfg.JWTSecretKey)
	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		UserID:      user.ID,
		AccessToken: accessToken,
	}, nil
}
