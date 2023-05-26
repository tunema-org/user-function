package clients

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tunema-org/user-function/internal/config"
)

func NewAWSSession(cfg *config.Config) (*session.Session, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSAccessKeyID,
			cfg.AWSSecretAccessKey,
			cfg.AWSSessionToken,
		),
	})
	if err != nil {
		return nil, err
	}

	return awsSession, nil
}

type S3 struct {
	sess   *session.Session
	client *s3.S3
	bucket string
}

func S3NewClient(sess *session.Session, bucket string) *S3 {
	return &S3{
		sess:   sess,
		bucket: bucket,
		client: s3.New(sess),
	}
}

func (s *S3) UploadFile(ctx context.Context, key string, file multipart.File) (*s3.PutObjectOutput, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   file,
	}

	return s.client.PutObjectWithContext(ctx, input)
}

func (s *S3) FindFile(ctx context.Context, bucket, key string) (*s3.GetObjectOutput, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	return s.client.GetObject(input)
}

func (s *S3) DeleteFile(ctx context.Context, bucket, key string) (*s3.DeleteObjectOutput, error) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	return s.client.DeleteObject(input)
}
