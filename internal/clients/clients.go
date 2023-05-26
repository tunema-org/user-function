package clients

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tunema-org/user-function/internal/config"
	"golang.org/x/sync/errgroup"
)

type Clients struct {
	DB      *pgxpool.Pool
	Session *session.Session
	S3      *S3
}

func New(ctx context.Context, cfg *config.Config) (*Clients, error) {
	var group errgroup.Group

	c := &Clients{}

	group.Go(func() error {
		session, err := NewAWSSession(cfg)
		if err != nil {
			return err
		}

		c.Session = session
		c.S3 = S3NewClient(session, cfg.S3Bucket)

		return nil
	})

	group.Go(func() error {
		var err error

		c.DB, err = NewPostgreSQLClient(ctx, cfg.DatabaseURL)
		if err != nil {
			return err
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}

	return c, nil
}
