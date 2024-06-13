package server

import (
	"encoding/json"
	"net/http"
	"server-bot/internal/functionalities"
	"time"
)

func (s *Server) AuthHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		functionalities.WriteJSON(w, http.StatusBadRequest, APIServerError{Error: "Invalid request"})
		return
	}

	user, err := s.db.GetUserByEmail(input.Email)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "Internal server error"})
		return
	}

	if user == nil {
		// Register new user
		user, err = s.db.RegisterUser(input.Email)
		if err != nil {
			functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "Failed to register user"})
			return
		}
	} else if time.Now().After(user.CodeValidUntil) {
		// Update verification code if expired
		err = s.db.UpdateUserVerificationCode(user)
		if err != nil {
			functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "Failed to update verification code"})
			return
		}
	}

	// Send the verification code via email (implementation dependent)
	functionalities.SendVerificationCodeEmail(user.Email, user.VerificationCode)
	functionalities.WriteJSON(w, http.StatusOK, APISuccessMessage{Message: "Verification code sent to your email."})
}

//
//
//

func (s *Server) AuthVerifyHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		functionalities.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	user, err := s.db.GetUserByEmail(input.Email)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		return
	}
	if user == nil || user.VerificationCode != input.Code || time.Now().After(user.CodeValidUntil) {
		functionalities.WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid or expired code"})
		return
	}

	// Code is correct and not expired, activate user
	user.IsActive = true
	token, tokenValidUntil := functionalities.GenerateTokenForUser(user)
	user.Token = token
	user.TokenValidUntil = tokenValidUntil
	err = s.db.ActivateUser(user)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to activate user"})
		return
	}

	// Send response with token
	functionalities.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "User verified and activated successfully",
		"error":   "",
		"token":   token,
	})
}
