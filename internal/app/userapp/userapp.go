package userapp

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/google/uuid"

	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
)

type UserProvider interface {
	SaveUser(ctx context.Context, u user.User) error
	GetUserByID(ctx context.Context, id user.UserID) (user.User, error)
	FindUserByName(ctx context.Context, name string) (user.User, error)
	FindUsersByIDs(ctx context.Context, ids []user.UserID) ([]user.User, error)
}

type App struct {
	up UserProvider
}

func New(up UserProvider) *App {
	return &App{
		up: up,
	}
}

func (a *App) CreateUser(ctx context.Context, u user.User) (user.UserID, error) {
	if err := validateEmail(u.Email); err != nil {
		return user.NilUserID(), err
	}
	if err := validateName(u.Nickname); err != nil {
		return user.NilUserID(), err
	}

	u.ID = generateUserID()

	if err := a.up.SaveUser(ctx, u); err != nil {
		return user.NilUserID(), err
	}

	return u.ID, nil
}

func (a *App) GetUserByID(ctx context.Context, id user.UserID) (user.User, error) {
	return a.up.GetUserByID(ctx, id)
}

func (a *App) FindUserByName(ctx context.Context, name string) (user.User, error) {
	return a.up.FindUserByName(ctx, name)
}

func generateUserID() user.UserID {
	return user.UserID(uuid.New())
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}

func validateName(name string) error {
	if len(name) > 100 {
		return fmt.Errorf("name is too long: %d > 100", len(name))
	}
	if len(name) < 2 {
		return fmt.Errorf("name is too short: %d < 2", len(name))
	}
	return nil
}
