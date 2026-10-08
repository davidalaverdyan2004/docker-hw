package httpserver

import (
	"employees/db"
	"encoding/json"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

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

	mtx.Lock()
	defer mtx.Unlock()

	emp.ID = nextID
	nextID++

	employees = append(employees, emp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)

}
