package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.healthHandler).Methods("GET")

	r.HandleFunc("/products", s.GetAllProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", s.GetProductHandler).Methods("GET")
	r.HandleFunc("/products", s.AddProductHandler).Methods("POST")           // S
	r.HandleFunc("/products/{id}", s.DeleteProductHandler).Methods("DELETE") // S
	r.HandleFunc("/products/{id}", s.UpdateProductHandler).Methods("PUT")    // S
	r.HandleFunc("/user", s.GetUserHandler).Methods("GET")                   // S
	r.HandleFunc("/user/{id}", s.GetAllUsersHandler).Methods("GET")          // S

	r.HandleFunc("/auth", s.AuthHandler).Methods("POST")
	r.HandleFunc("/verify", s.AuthVerifyHandler).Methods("POST")
	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
