package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserID uuid.UUID

type User struct {
	ID        UserID
	Nickname  string
	Email     string
	CreatedAt time.Time
}

func NilUserID() UserID {
	return UserID(uuid.Nil)
}
