package router

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/stepan2volkov/user-group-srv/internal/api/openapi"
	"github.com/stepan2volkov/user-group-srv/internal/entities/group"
)

// CreateGroup implements openapi.ServerInterface
func (s *Router) CreateGroup(w http.ResponseWriter, r *http.Request) {
	g := openapi.Group{}
	err := json.NewDecoder(r.Body).Decode(&g)
	defer r.Body.Close()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	groupType, err := group.GetGroupType(string(*g.GroupType))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	groupID, err := s.groupApp.CreateGroup(r.Context(), group.Group{
		Title: g.Title,
		Type:  groupType,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(openapi.CreatedGroupResp{
		Id: [16]byte(groupID),
	}); err != nil {
		log.Printf("error when encoding created group: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// FindGroupByTitle implements openapi.ServerInterface
func (s *Router) FindGroupByTitle(w http.ResponseWriter, r *http.Request, title string) {
	g, err := s.groupApp.FindGroupByTitle(r.Context(), title)
	if errors.Is(err, group.ErrGroupNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(g); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetGroupByID implements openapi.ServerInterface
func (s *Router) GetGroupByID(w http.ResponseWriter, r *http.Request, groupID uuid.UUID) {
	g, err := s.groupApp.GetGroupByID(r.Context(), group.GroupID(groupID))
	if errors.Is(err, group.ErrGroupNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(g); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
