package config

import (
	"fmt"
	"time"
)

type Config struct {
	Address      string
	JWTSecretKey string
	JWTDuration  time.Duration
	DatabaseURL  string
}

func New() *Config {
	var c Config

	port := LookupEnv("PORT", 5000)
	c = Config{
		Address:      LookupEnv("APP_ADDRESS", fmt.Sprintf(":%d", port)),
		JWTSecretKey: LookupEnv("JWT_SECRET_KEY", "secret"),
		JWTDuration: LookupEnv("JWT_DURATION", time.Duration(
			time.Now().Add(time.Hour*24*30).Unix())),
		DatabaseURL: LookupEnv("DATABASE_URL", ""),
	}

	return &c
}
