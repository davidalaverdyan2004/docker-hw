package httpserver

import (
	"employees/db"
	"encoding/json"
	"net/http"
)

func (h *Handler)CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var emp db.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "err JSON", http.StatusBadRequest)
		return
	}

	if emp.FullName == "" {
		http.Error(w, "пустая строка", http.StatusBadRequest)
		return
	}


	emp, err = db.AddEmployee(r.Context(), h.pool, emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)

}
