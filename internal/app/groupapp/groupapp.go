package groupapp

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
)

func generateGroupID() group.GroupID {
	return group.GroupID(uuid.New())
}

type GroupProvider interface {
	SaveGroup(ctx context.Context, g group.Group) error
	GetGroupByID(ctx context.Context, id group.GroupID) (group.Group, error)
	FindGroupByTitle(ctx context.Context, title string) (group.Group, error)
	FindGroupsByIDs(ctx context.Context, id []group.GroupID) ([]group.Group, error)
}

type App struct {
	gp GroupProvider
}

func New(gp GroupProvider) *App {
	return &App{
		gp: gp,
	}
}

func (a *App) CreateGroup(ctx context.Context, g group.Group) (group.GroupID, error) {
	if err := validateTitle(g.Title); err != nil {
		return group.NilGroupID(), err
	}

	g.ID = generateGroupID()

	if err := a.gp.SaveGroup(ctx, g); err != nil {
		return group.NilGroupID(), err
	}

	return g.ID, nil
}

func (a *App) GetGroupByID(ctx context.Context, id group.GroupID) (group.Group, error) {
	return a.gp.GetGroupByID(ctx, id)
}

func (a *App) FindGroupByTitle(ctx context.Context, title string) (group.Group, error) {
	return a.gp.FindGroupByTitle(ctx, title)
}

func validateTitle(title string) error {
	if len(title) > 200 {
		return fmt.Errorf("title is too long: %d > 200", len(title))
	}
	if len(title) < 2 {
		return fmt.Errorf("title is too short: %d < 2", len(title))
	}
	return nil
}
