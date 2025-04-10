package repository

import (
	"context"
	"database/sql"
	"errors"

	"rockstar-corp/internal/domain/user"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, u *user.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, u.Name, u.Email, u.Password).Scan(&u.ID)
	return err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*user.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
