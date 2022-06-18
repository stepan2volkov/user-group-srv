package group

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrGroupNotFound = errors.New("group not found")
)

type GroupID uuid.UUID

type GroupType string

const (
	Project        GroupType = "project"
	Organization   GroupType = "org"
	CorporateGroup GroupType = "corporate_group"
	Community      GroupType = "community"
)

type Group struct {
	ID        GroupID
	Title     string
	Type      GroupType
	CreatedAt time.Time
}

func GetGroupType(v string) (GroupType, error) {
	switch GroupType(v) {
	case Project:
		return Project, nil
	case Organization:
		return Organization, nil
	case CorporateGroup:
		return CorporateGroup, nil
	case Community:
		return Community, nil
	default:
		return "", fmt.Errorf("unknown group-type '%s'", v)
	}
}
