package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"server-bot/internal/functionalities"
	"server-bot/internal/models"
	"strconv"
)

func (s *Server) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	prods, err := s.db.GetAllProducts()
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	functionalities.WriteJSON(w, http.StatusOK, prods)
}

func (s *Server) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		functionalities.WriteJSON(w, http.StatusBadRequest, APIServerError{Error: err.Error()})
		return
	}

	prod, err := s.db.GetProduct(uint(id))
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	functionalities.WriteJSON(w, http.StatusOK, prod)
}

func (s *Server) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//
	// Retrieve the password from the query parameters
	providedPassword := r.URL.Query().Get("p")

	// Check if the provided password matches the predefined password
	if providedPassword != os.Getenv("ACCESS_PASSWORD") {
		functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Unauthorized"})
		return
	}
	//

	newProd := new(models.Product)

	if err := json.NewDecoder(r.Body).Decode(newProd); err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	err = s.db.AddProduct(newProd)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	functionalities.WriteJSON(w, http.StatusOK, APISuccessMessage{Message: "Product added successfully!"})
}
