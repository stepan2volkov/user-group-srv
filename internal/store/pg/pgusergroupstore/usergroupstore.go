package pgusergroupstore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
	"github.com/stepan2volkov/user-group-srv/internal/entities/usergroup"
)

var _ usergroupapp.UserGroupProvider = &UserGroupMapper{}

type UserGroupMapper struct {
	db *sql.DB
}

func New(db *sql.DB) *UserGroupMapper {
	return &UserGroupMapper{
		db: db,
	}
}

// AddUserToGroup implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) AddUserToGroup(ctx context.Context, ug usergroup.UserGroup) error {
	_, err := ugm.db.ExecContext(ctx,
		`INSERT INTO usergroups(user_id, group_id)
		VALUES ($1, $2)`, ug.UserID, ug.GroupID)

	return err
}

// DropUserFromGroup implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) DropUserFromGroup(ctx context.Context, ug usergroup.UserGroup) error {
	_, err := ugm.db.ExecContext(ctx,
		`DROP FROM usergroups
		WHERE user_id = $1 AND group_id = $2`, ug.UserID, ug.GroupID)

	return err
}

// FindGroupIDsByUserID implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) FindGroupIDsByUserID(ctx context.Context, id user.UserID) ([]group.GroupID, error) {
	ret := make([]group.GroupID, 0, 10)

	rows, err := ugm.db.QueryContext(ctx,
		`SELECT group_id 
		FROM usergroups 
		WHERE user_id = $1`, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, usergroup.ErrNoAssignedGroup
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id group.GroupID
		if err = rows.Scan(id); err != nil {
			return nil, err
		}
		ret = append(ret, id)
	}

	return ret, nil
}

// FindUserIDsByGroupID implements usergroupapp.UserGroupProvider
func (ugm *UserGroupMapper) FindUserIDsByGroupID(ctx context.Context, id group.GroupID) ([]user.UserID, error) {
	ret := make([]user.UserID, 0, 10)

	rows, err := ugm.db.QueryContext(ctx,
		`SELECT user_id 
		FROM usergroups 
		WHERE group_id = $1`, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, usergroup.ErrNoUsersInGroup
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id user.UserID
		if err = rows.Scan(id); err != nil {
			return nil, err
		}
		ret = append(ret, id)
	}

	return ret, nil
}
