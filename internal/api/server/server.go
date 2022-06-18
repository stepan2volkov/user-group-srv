package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/stepan2volkov/user-group-srv/internal/api/openapi"
	"github.com/stepan2volkov/user-group-srv/internal/app/userapp"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
)

var _ openapi.ServerInterface = &Server{}

type Server struct {
	userApp *userapp.App
}

func New(userApp *userapp.App) *Server {
	return &Server{
		userApp: userApp,
	}
}

// CreateUser implements openapi.ServerInterface
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := openapi.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = s.userApp.CreateUser(r.Context(), user.User{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AddUserToGroup implements openapi.ServerInterface
func (s *Server) AddUserToGroup(w http.ResponseWriter, r *http.Request) {

}

// CreateGroup implements openapi.ServerInterface
func (s *Server) CreateGroup(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// DropUserFromGroup implements openapi.ServerInterface
func (s *Server) DropUserFromGroup(w http.ResponseWriter, r *http.Request, userID uuid.UUID, groupID uuid.UUID) {
	panic("unimplemented")
}

// FindGroupByTitle implements openapi.ServerInterface
func (s *Server) FindGroupByTitle(w http.ResponseWriter, r *http.Request, title string) {
	panic("unimplemented")
}

// FindGroupsByUserID implements openapi.ServerInterface
func (s *Server) FindGroupsByUserID(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	panic("unimplemented")
}

// FindUserByNickname implements openapi.ServerInterface
func (s *Server) FindUserByNickname(w http.ResponseWriter, r *http.Request, nickname string) {
	panic("unimplemented")
}

// FindUsersByGroupID implements openapi.ServerInterface
func (s *Server) FindUsersByGroupID(w http.ResponseWriter, r *http.Request, groupID uuid.UUID) {
	panic("unimplemented")
}

// GetGroupByID implements openapi.ServerInterface
func (s *Server) GetGroupByID(w http.ResponseWriter, r *http.Request, groupID uuid.UUID) {
	panic("unimplemented")
}

// GetUserByID implements openapi.ServerInterface
func (s *Server) GetUserByID(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	panic("unimplemented")
}
