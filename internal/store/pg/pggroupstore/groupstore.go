package pggroupstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/stepan2volkov/user-group-srv/internal/app/groupapp"
	"github.com/stepan2volkov/user-group-srv/internal/app/usergroupapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
)

var _ groupapp.GroupProvider = &GroupMapper{}
var _ usergroupapp.GroupProvider = &GroupMapper{}

type GroupMapper struct {
	db *sql.DB
}

func New(db *sql.DB) *GroupMapper {
	return &GroupMapper{
		db: db,
	}
}

// SaveUser implements app.UserProvider
func (um *GroupMapper) SaveGroup(ctx context.Context, g group.Group) error {
	_, err := um.db.ExecContext(ctx,
		`INSERT INTO groups(id, title, group_type)
		VALUES ($1, $2, $3)`, g.ID, g.Title, g.Type)

	return err
}

// GetUserByID implements usergroupapp.UserProvider
func (um *GroupMapper) GetGroupByID(ctx context.Context, id group.GroupID) (group.Group, error) {
	return um.findGroup(ctx, "id = $1", id)
}

// FindUsersByIDs implements usergroupapp.UserProvider
func (um *GroupMapper) FindGroupsByIDs(ctx context.Context, ids []group.GroupID) ([]group.Group, error) {
	uuids := &pgtype.UUIDArray{}
	uuids.Set(ids)
	return um.findGroups(ctx, "id = ANY($1)", uuids)
}

// FindUserByName implements app.UserProvider
func (um *GroupMapper) FindGroupByTitle(ctx context.Context,
	name string) (group.Group, error) {

	return um.findGroup(ctx, "lower(name) like $1", strings.ToLower(name))
}

func (um *GroupMapper) findGroup(ctx context.Context, where string,
	v ...interface{}) (group.Group, error) {

	g := group.Group{}
	var id uuid.UUID

	err := um.db.QueryRowContext(ctx, fmt.Sprintf(`
			SELECT id, title, group_type, created_at
			FROM groups WHERE %s`, where), v...).Scan(
		&id, &g.Title, &g.Type, &g.CreatedAt,
	)
	g.ID = group.GroupID(id)

	// We don't have to depend only on sql db
	if errors.Is(err, sql.ErrNoRows) {
		return group.Group{}, group.ErrGroupNotFound
	}
	if err != nil {
		return group.Group{}, err
	}

	return g, nil
}

func (um *GroupMapper) findGroups(ctx context.Context, where string,
	v ...interface{}) ([]group.Group, error) {

	ret := make([]group.Group, 0, 1)

	rows, err := um.db.QueryContext(ctx, fmt.Sprintf(`
			SELECT id, title, group_type, created_at
			FROM groups WHERE %s`, where), v...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, group.ErrGroupNotFound
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		g := group.Group{}
		var id uuid.UUID
		err := rows.Scan(&id, &g.Title, &g.Type, &g.CreatedAt)
		if err != nil {
			return nil, err
		}
		g.ID = group.GroupID(id)

		ret = append(ret, g)
	}

	return ret, nil
}
