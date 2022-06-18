package usergroup

import (
	"errors"

	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
)

var (
	ErrNoAssignedGroup = errors.New("user hasn't assigned to any group")
	ErrNoUsersInGroup  = errors.New("group doesn't have any user")
)

type UserGroup struct {
	UserID  user.UserID
	GroupID group.GroupID
}
