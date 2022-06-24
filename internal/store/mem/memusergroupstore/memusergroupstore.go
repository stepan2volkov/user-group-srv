package memusergroupstore

import (
	"context"
	"database/sql"
	"errors"
	"sync"

	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
	"github.com/stepan2volkov/user-group-srv/internal/entities/usergroup"
)

var _ usergroupapp.UserGroupProvider = &UserGroupMapper{}

type UserGroupMapper struct {
	m          sync.Mutex
	usergroups map[usergroup.UserGroup]struct{}
}

func New() *UserGroupMapper {
	return &UserGroupMapper{
		m:          sync.Mutex{},
		usergroups: make(map[usergroup.UserGroup]struct{}),
	}
}

// AddUserToGroup implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) AddUserToGroup(ctx context.Context, ug usergroup.UserGroup) error {
	ugm.m.Lock()
	defer ugm.m.Unlock()

	if _, found := ugm.usergroups[ug]; found {
		return errors.New("user already in group")
	}

	ugm.usergroups[ug] = struct{}{}
	return nil
}

// DropUserFromGroup implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) DropUserFromGroup(ctx context.Context, ug usergroup.UserGroup) error {
	ugm.m.Lock()
	defer ugm.m.Unlock()

	delete(ugm.usergroups, ug)

	return nil
}

// FindGroupIDsByUserID implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) FindGroupIDsByUserID(ctx context.Context, id user.UserID) ([]group.GroupID, error) {
	ugm.m.Lock()
	defer ugm.m.Unlock()

	ret := make([]group.GroupID, 0, 1)

	for ug := range ugm.usergroups {
		if ug.UserID == id {
			ret = append(ret, ug.GroupID)
		}
	}

	if len(ret) == 0 {
		return nil, sql.ErrNoRows
	}

	return ret, nil
}

// FindUserIDsByGroupID implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) FindUserIDsByGroupID(ctx context.Context, id group.GroupID) ([]user.UserID, error) {
	ugm.m.Lock()
	defer ugm.m.Unlock()

	ret := make([]user.UserID, 0, 1)

	for ug := range ugm.usergroups {
		if ug.GroupID == id {
			ret = append(ret, ug.UserID)
		}
	}

	if len(ret) == 0 {
		return nil, sql.ErrNoRows
	}

	return ret, nil
}
