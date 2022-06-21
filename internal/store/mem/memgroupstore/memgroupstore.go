package memgroupstore

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/stepan2volkov/user-group-srv/internal/app/groupapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
)

var _ groupapp.GroupProvider = &GroupMapper{}
var _ usergroupapp.GroupProvider = &GroupMapper{}

type GroupMapper struct {
	m             sync.Mutex
	groupsByID    map[group.GroupID]group.Group
	groupsByTitle map[string]group.Group
}

func New() *GroupMapper {
	return &GroupMapper{
		m:             sync.Mutex{},
		groupsByID:    make(map[group.GroupID]group.Group),
		groupsByTitle: make(map[string]group.Group),
	}
}

// SaveUser implements app.UserProvider
func (um *GroupMapper) SaveGroup(ctx context.Context, g group.Group) error {
	um.m.Lock()
	defer um.m.Unlock()

	if _, found := um.groupsByTitle[g.Title]; found {
		return fmt.Errorf("user with title %s has already existed", g.Title)
	}

	g.CreatedAt = time.Now()
	um.groupsByID[g.ID] = g
	um.groupsByTitle[g.Title] = g

	return nil
}

// GetUserByID implements usergroupapp.UserProvider
func (um *GroupMapper) GetGroupByID(ctx context.Context, id group.GroupID) (group.Group, error) {
	um.m.Lock()
	defer um.m.Unlock()

	if g, found := um.groupsByID[id]; found {
		return g, nil
	}

	return group.Group{}, sql.ErrNoRows
}

// FindUsersByIDs implements usergroupapp.UserProvider
func (um *GroupMapper) FindGroupsByIDs(ctx context.Context, ids []group.GroupID) ([]group.Group, error) {
	um.m.Lock()
	defer um.m.Unlock()

	ret := make([]group.Group, 0, len(ids))

	for _, id := range ids {
		if g, found := um.groupsByID[id]; found {
			ret = append(ret, g)
		}
	}

	if len(ret) == 0 {
		return nil, sql.ErrNoRows
	}

	return ret, nil
}

// FindUserByName implements app.UserProvider
func (um *GroupMapper) FindGroupByTitle(ctx context.Context,
	name string) (group.Group, error) {

	um.m.Lock()
	defer um.m.Unlock()

	if g, found := um.groupsByTitle[name]; found {
		return g, nil
	}

	return group.Group{}, sql.ErrNoRows
}
