package memuserstore

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/stepan2volkov/user-group-srv/internal/app/userapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
)

var _ userapp.UserProvider = &UserMapper{}
var _ usergroupapp.UserProvider = &UserMapper{}

type UserMapper struct {
	m           sync.Mutex
	usersByID   map[user.UserID]user.User
	usersByName map[string]user.User
}

func New() *UserMapper {
	return &UserMapper{
		m:           sync.Mutex{},
		usersByID:   make(map[user.UserID]user.User),
		usersByName: make(map[string]user.User),
	}
}

// SaveUser implements app.UserProvider
func (um *UserMapper) SaveUser(ctx context.Context, u user.User) error {
	um.m.Lock()
	defer um.m.Unlock()

	// There is no constraint about email, only nickname
	if _, found := um.usersByName[u.Nickname]; found {
		return errors.New("nickname has already existed")
	}

	um.usersByID[u.ID] = u
	um.usersByName[u.Nickname] = u

	return nil
}

// GetUserByID implements usergroupapp.UserProvider
func (um *UserMapper) GetUserByID(ctx context.Context, id user.UserID) (user.User, error) {
	um.m.Lock()
	defer um.m.Unlock()

	if u, found := um.usersByID[id]; found {
		return u, nil
	}

	return user.User{}, sql.ErrNoRows
}

// FindUsersByIDs implements usergroupapp.UserProvider
func (um *UserMapper) FindUsersByIDs(ctx context.Context, ids []user.UserID) ([]user.User, error) {
	um.m.Lock()
	defer um.m.Unlock()

	ret := make([]user.User, 0, len(ids))

	for _, id := range ids {
		if u, found := um.usersByID[id]; found {
			ret = append(ret, u)
		}
	}

	if len(ids) == 0 {
		return nil, sql.ErrNoRows
	}

	return ret, nil
}

// FindUserByName implements app.UserProvider
func (um *UserMapper) FindUserByName(ctx context.Context,
	name string) (user.User, error) {

	um.m.Lock()
	defer um.m.Unlock()

	if u, found := um.usersByName[name]; found {
		return u, nil
	}

	return user.User{}, sql.ErrNoRows
}
