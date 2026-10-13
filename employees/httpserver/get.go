package httpserver

import (
	"employees/db"
	"encoding/json"
	"net/http"
)

func (h *Handler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := db.GetEmployees(r.Context(), h.pool)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)

}
