package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server-bot/internal/database"
	"server-bot/internal/functionalities"
	"strings"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	main := r.PathPrefix("/api").Subrouter()
	main.Use(Logger)

	main.HandleFunc("/", s.healthHandler).Methods("GET")                 // health
	main.HandleFunc("/products", s.GetAllProductsHandler).Methods("GET") // products
	main.HandleFunc("/products/{id}", s.GetProductHandler).Methods("GET")
	main.HandleFunc("/auth", s.AuthHandler).Methods("POST") // authorization
	main.HandleFunc("/verify", s.AuthVerifyHandler).Methods("POST")

	secure := main.PathPrefix("").Subrouter()
	secure.Use(SecureMW(s.db))
	secure.HandleFunc("/products", s.AddProductHandler).Methods("POST")
	secure.HandleFunc("/products/{id}", s.DeleteProductHandler).Methods("DELETE")
	secure.HandleFunc("/products/{id}", s.UpdateProductHandler).Methods("PUT")
	secure.HandleFunc("/user", s.GetAllUsersHandler).Methods("GET")
	secure.HandleFunc("/user/{id}", s.GetUserHandler).Methods("GET")

	return main
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s ~ \"%s\"", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func SecureMW(s database.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				// prod:
				//functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Invalid authorization token"})
				functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Missing Authorization header"})
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Token" {
				// prod:
				//functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Invalid authorization token"})
				functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Invalid Authorization header format"})
				return
			}

			token := parts[1]

			// Verify the token
			user, err := s.GetUserByToken(token)
			if err != nil {
				functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "Internal server error"})
				return
			}
			if user == nil {
				// prod:
				//functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Invalid authorization token"})
				functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Invalid token"})
				return
			}

			// Add user information to the request context if needed
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
