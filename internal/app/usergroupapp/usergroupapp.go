package usergroupapp

import (
	"context"
	"fmt"

	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
	"github.com/stepan2volkov/user-group-srv/internal/entities/usergroup"
)

// UserProvider can be db or wrapper of rpc service.
type UserProvider interface {
	GetUserByID(ctx context.Context, id user.UserID) (user.User, error)
	FindUsersByIDs(ctx context.Context, ids []user.UserID) ([]user.User, error)
}

// GroupProvider can be db or wrapper of rpc service
type GroupProvider interface {
	GetGroupByID(ctx context.Context,
		id group.GroupID) (group.Group, error)
	FindGroupsByIDs(ctx context.Context,
		id []group.GroupID) ([]group.Group, error)
}

type UserGroupProvider interface {
	AddUserToGroup(ctx context.Context, ug usergroup.UserGroup) error
	DropUserFromGroup(ctx context.Context, ug usergroup.UserGroup) error
	FindUserIDsByGroupID(ctx context.Context, id group.GroupID) ([]user.UserID, error)
	FindGroupIDsByUserID(ctx context.Context, id user.UserID) ([]group.GroupID, error)
}

type App struct {
	up  UserProvider
	gp  GroupProvider
	ugp UserGroupProvider
}

func New(up UserProvider,
	gp GroupProvider,
	ugp UserGroupProvider) *App {
	return &App{
		up:  up,
		gp:  gp,
		ugp: ugp,
	}
}

func (a *App) AddUserToGroup(
	ctx context.Context,
	ug usergroup.UserGroup,
) error {

	// TODO: Проверка, что пользователь и группа существуют
	// TODO: Проверка, что пользователь не назначен в группу
	// TODO: Назначение группы пользователю

	_, err := a.gp.GetGroupByID(ctx, ug.GroupID)
	if err != nil {
		return fmt.Errorf("error when searching group: %w", err)
	}

	_, err = a.up.GetUserByID(ctx, user.UserID(ug.GroupID))
	if err != nil {
		return fmt.Errorf("error when searching user: %w", err)
	}

	return a.ugp.AddUserToGroup(ctx, ug)
}

func (a *App) DropUserFromGroup(
	ctx context.Context,
	ug usergroup.UserGroup,
) error {
	// TODO: Проверка, что пользователь принадлежит группе
	// TODO: Удаление пользователя из группы

	_, err := a.gp.GetGroupByID(ctx, ug.GroupID)
	if err != nil {
		return fmt.Errorf("error when searching group: %w", err)
	}

	_, err = a.up.GetUserByID(ctx, ug.UserID)
	if err != nil {
		return fmt.Errorf("error when searching user: %w", err)
	}

	return a.ugp.DropUserFromGroup(ctx, ug)
}

func (a *App) FindUsersByGroupID(
	ctx context.Context,
	groupID group.GroupID,
) ([]user.User, error) {

	userIDs, err := a.ugp.FindUserIDsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	users, err := a.up.FindUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *App) FindGroupsByUserID(
	ctx context.Context,
	userID user.UserID,
) ([]group.Group, error) {

	groupIDs, err := a.ugp.FindGroupIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	groups, err := a.gp.FindGroupsByIDs(ctx, groupIDs)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
