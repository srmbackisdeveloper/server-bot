package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"server-bot/internal/functionalities"
	"strconv"
)

func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := s.db.GetUser(uint(userId))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	functionalities.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := s.db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	functionalities.WriteJSON(w, http.StatusOK, users)
}
