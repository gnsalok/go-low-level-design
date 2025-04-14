package http

import (
	"ddd-go/internal/domain/services"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user services.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userService.RegisterUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a user by ID will go here
	// This is a placeholder for demonstration purposes
	w.WriteHeader(http.StatusNotImplemented)
}
