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

	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSSessionToken    string
	AWSRegion          string

	S3Bucket string
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

		AWSAccessKeyID:     LookupEnv("AWS_ACCESS_KEY_ID", ""),
		AWSSecretAccessKey: LookupEnv("AWS_SECRET_ACCESS_KEY", ""),
		AWSSessionToken:    LookupEnv("AWS_SESSION_TOKEN", ""),
		AWSRegion:          LookupEnv("AWS_REGION", "us-east-1"),

		S3Bucket: LookupEnv("S3_BUCKET", ""),
	}

	return &c
}
