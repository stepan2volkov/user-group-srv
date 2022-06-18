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
	FindUserByName(ctx context.Context, name string) (user.User, error)
}

type App struct {
	up UserProvider
}

func New(up UserProvider) *App {
	return &App{
		up: up,
	}
}

func (a *App) CreateUser(ctx context.Context, u user.User) error {
	if err := validateEmail(u.Email); err != nil {
		return err
	}
	if err := validateName(u.Nickname); err != nil {
		return err
	}

	u.ID = generateUserID()

	return a.up.SaveUser(ctx, u)
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
