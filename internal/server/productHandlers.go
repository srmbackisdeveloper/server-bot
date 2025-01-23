package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

	newProd := new(models.Product)

	if err := json.NewDecoder(r.Body).Decode(newProd); err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	err := s.db.AddProduct(newProd)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	functionalities.WriteJSON(w, http.StatusCreated, APISuccessMessage{Message: "Product added successfully!"})
}

func (s *Server) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "invalid id"})
		return
	}

	err = s.db.DeleteProduct(uint(id))
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	functionalities.WriteJSON(w, http.StatusNoContent, APISuccessMessage{Message: "Product deleted successfully!"})
}

func (s *Server) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	providedPassword := r.URL.Query().Get("p")

	if providedPassword != os.Getenv("ACCESS_PASSWORD") {
		functionalities.WriteJSON(w, http.StatusUnauthorized, APIServerError{Error: "Unauthorized"})
		return
	}

	idString := mux.Vars(r)["id"] // get ID from URL
	id, err := strconv.Atoi(idString)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "Invalid ID format"})
		return
	}

	var updateProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updateProduct); err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: "Error decoding request body"})
		return
	}

	err = s.db.UpdateProduct(uint(id), &updateProduct)
	if err != nil {
		functionalities.WriteJSON(w, http.StatusInternalServerError, APIServerError{Error: err.Error()})
		return
	}

	functionalities.WriteJSON(w, http.StatusOK, APISuccessMessage{Message: "Product updated successfully!"})
}
