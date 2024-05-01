package functionalities

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, anything interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(anything)
	if err != nil {
		return
	}
}
