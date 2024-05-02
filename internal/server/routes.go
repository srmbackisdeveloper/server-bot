package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("", s.healthHandler).Methods("GET")

	r.HandleFunc("/products", s.GetAllProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", s.GetProductHandler).Methods("GET")
	r.HandleFunc("/products", s.AddProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", s.DeleteProductHandler).Methods("DELETE")
	r.HandleFunc("/products/{id}", s.UpdateProductHandler).Methods("PUT")

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}
	_, _ = w.Write(jsonResp)
}
