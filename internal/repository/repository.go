package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertUser(ctx context.Context, user User) (int, error) {
	existsQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	var alreadyExists bool
	err := r.db.QueryRow(ctx, existsQuery, user.Email).Scan(&alreadyExists)
	if err != nil {
		return -1, err
	}

	if alreadyExists {
		return -1, ErrUserAlreadyExists
	}

	insertQuery := `INSERT INTO users (username, email, password, profile_img_url) VALUES ($1, $2, $3, $4) RETURNING id`
	var userID int

	err = r.db.QueryRow(ctx, insertQuery, user.Username, user.Email, user.Password, user.ProfileImgURL).
		Scan(&userID)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (r *Repository) FindUserByID(ctx context.Context, id int) (User, error) {
	query := `SELECT id, username, email, password, profile_img_url FROM users WHERE id = $1`
	var user User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.ProfileImgURL)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (User, error) {
	query := `SELECT id, username, email, password, profile_img_url FROM users WHERE email = $1`
	var user User
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.ProfileImgURL)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
