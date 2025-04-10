package user

import (
	"context"
	"errors"
)

type UseCase interface {
	Register(ctx context.Context, u *User) error
	GetProfile(ctx context.Context, id int64) (*User, error)
}

type useCase struct {
	repo Repository
}

func New(repo Repository) UseCase {
	return &useCase{repo: repo}
}

func (uc *useCase) Register(ctx context.Context, u *User) error {
	existing, _ := uc.repo.GetByEmail(ctx, u.Email)
	if existing != nil {
		return errors.New("user already exists")
	}

	// TODO: hash password before storing (bcrypt)
	return uc.repo.Create(ctx, u)
}

func (uc *useCase) GetProfile(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetByID(ctx, id)
}
