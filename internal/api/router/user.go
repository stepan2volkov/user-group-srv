package router

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/stepan2volkov/user-group-srv/internal/api/openapi"
	"github.com/stepan2volkov/user-group-srv/internal/entities/user"
)

// CreateUser implements openapi.ServerInterface
func (s *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := openapi.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	defer r.Body.Close()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := s.userApp.CreateUser(r.Context(), user.User{
		Nickname: u.Nickname,
		Email:    u.Email,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&openapi.CreatedUserResp{
		Id: [16]byte(userID),
	}); err != nil {
		log.Printf("error when encoding created user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Router) FindUserByNickname(w http.ResponseWriter, r *http.Request, nickname string) {
	u, err := s.userApp.FindUserByName(r.Context(), nickname)
	if errors.Is(err, user.ErrUserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(u); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetUserByID implements openapi.ServerInterface
func (s *Router) GetUserByID(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	u, err := s.userApp.GetUserByID(r.Context(), user.UserID(userID))
	if errors.Is(err, user.ErrUserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(u); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
