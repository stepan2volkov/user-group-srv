package router

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stepan2volkov/user-group-srv/internal/api/openapi"
	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
	"github.com/stepan2volkov/user-group-srv/internal/entities/usergroup"
)

// AddUserToGroup implements openapi.ServerInterface
func (s *Router) AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	ug := openapi.UserGroup{}
	err := json.NewDecoder(r.Body).Decode(&ug)
	defer r.Body.Close()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = s.userGroupApp.AddUserToGroup(r.Context(), usergroup.UserGroup{
		UserID:  user.UserID(ug.UserId),
		GroupID: group.GroupID(ug.GroupId),
	}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DropUserFromGroup implements openapi.ServerInterface
func (s *Router) DropUserFromGroup(w http.ResponseWriter, r *http.Request, userID uuid.UUID, groupID uuid.UUID) {
	if err := s.userGroupApp.DropUserFromGroup(r.Context(), usergroup.UserGroup{
		UserID:  user.UserID(userID),
		GroupID: group.GroupID(groupID),
	}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// FindGroupsByUserID implements openapi.ServerInterface
func (s *Router) FindGroupsByUserID(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	groups, err := s.userGroupApp.FindGroupsByUserID(r.Context(), user.UserID(userID))
	if errors.Is(err, usergroup.ErrNoAssignedGroup) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ret := make([]openapi.Group, 0, len(groups))
	for _, g := range groups {
		formattedTime := g.CreatedAt.Format(time.RFC3339)
		groupType := openapi.GroupGroupType(g.Type)
		ret = append(ret, openapi.Group{
			CreatedAt: &formattedTime,
			GroupType: &groupType,
			Id:        (*uuid.UUID)(&g.ID),
			Title:     g.Title,
		})
	}

	if err = json.NewEncoder(w).Encode(ret); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// FindUsersByGroupID implements openapi.ServerInterface
func (s *Router) FindUsersByGroupID(w http.ResponseWriter, r *http.Request, groupID uuid.UUID) {
	users, err := s.userGroupApp.FindUsersByGroupID(r.Context(), group.GroupID(groupID))
	if errors.Is(err, usergroup.ErrNoAssignedGroup) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret := make([]openapi.User, 0, len(users))
	for _, u := range users {
		formattedTime := u.CreatedAt.Format(time.RFC3339)
		ret = append(ret, openapi.User{
			CreatedAt: &formattedTime,
			Email:     u.Email,
			Id:        (*uuid.UUID)(&u.ID),
			Nickname:  u.Nickname,
		})
	}

	if err = json.NewEncoder(w).Encode(ret); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
