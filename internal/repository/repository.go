package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tunema-org/user-function/model"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	existsQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	var alreadyExists bool
	err := r.db.QueryRow(ctx, existsQuery, user.Email).Scan(&alreadyExists)
	if err != nil {
		return model.User{}, err
	}

	if alreadyExists {
		return model.User{}, ErrUserAlreadyExists
	}

	insertQuery := `INSERT INTO users (username, email, password, profile_img_url) VALUES ($1, $2, $3, $4) RETURNING id`
	var userID int

	err = r.db.QueryRow(ctx, insertQuery, user.Username, user.Email, user.Password, user.ProfileImgURL).
		Scan(&userID)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userID

	return user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, id int, user model.User) error {
	query := `UPDATE users SET username = $1, profile_img_url = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, user.Username, user.ProfileImgURL, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindUserByID(ctx context.Context, id int) (model.User, error) {
	query := `SELECT id, username, email, password, profile_img_url FROM users WHERE id = $1`
	var user model.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.ProfileImgURL)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	query := `SELECT id, username, email, password, profile_img_url FROM users WHERE email = $1`
	var user model.User
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.ProfileImgURL)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
