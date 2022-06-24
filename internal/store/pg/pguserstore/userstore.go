package pguserstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/stepan2volkov/user-group-srv/internal/app/userapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
)

var _ userapp.UserProvider = &UserMapper{}
var _ usergroupapp.UserProvider = &UserMapper{}

type UserMapper struct {
	db *sql.DB
}

func New(db *sql.DB) *UserMapper {
	return &UserMapper{
		db: db,
	}
}

// SaveUser implements app.UserProvider
func (um *UserMapper) SaveUser(ctx context.Context, u user.User) error {
	_, err := um.db.ExecContext(ctx,
		`INSERT INTO users(id, nickname, email)
		VALUES ($1, $2, $3)`, u.ID, u.Nickname, u.Email)

	return err
}

// GetUserByID implements usergroupapp.UserProvider
func (um *UserMapper) GetUserByID(ctx context.Context, id user.UserID) (user.User, error) {
	return um.findUser(ctx, "id = $1", id)
}

// FindUsersByIDs implements usergroupapp.UserProvider
func (um *UserMapper) FindUsersByIDs(ctx context.Context, ids []user.UserID) ([]user.User, error) {
	uuids := &pgtype.UUIDArray{}
	uuids.Set(ids)
	return um.findUsers(ctx, "id = ANY($1)", uuids)
}

// FindUserByName implements app.UserProvider
func (um *UserMapper) FindUserByName(ctx context.Context,
	name string) (user.User, error) {

	return um.findUser(ctx, "lower(name) like $1", strings.ToLower(name))
}

func (um *UserMapper) findUser(ctx context.Context, where string,
	v ...interface{}) (user.User, error) {

	u := user.User{}
	var id uuid.UUID

	err := um.db.QueryRowContext(ctx, fmt.Sprintf(`
			SELECT id, nickname, email, created_at
			FROM users WHERE %s`, where), v...).Scan(
		&id, &u.Nickname, &u.Email, &u.CreatedAt,
	)

	u.ID = user.UserID(id)

	// We don't have to depend only on sql db
	if errors.Is(err, sql.ErrNoRows) {
		return user.User{}, user.ErrUserNotFound
	}
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (um *UserMapper) findUsers(ctx context.Context, where string,
	v ...interface{}) ([]user.User, error) {

	ret := make([]user.User, 0, 1)

	rows, err := um.db.QueryContext(ctx, fmt.Sprintf(`
			SELECT id, nick_name, email, created_at
			FROM users WHERE %s`, where), v...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, user.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		u := user.User{}
		err := rows.Scan(&id, &u.Nickname, &u.Email, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		u.ID = user.UserID(id)

		ret = append(ret, u)
	}

	return ret, nil
}
