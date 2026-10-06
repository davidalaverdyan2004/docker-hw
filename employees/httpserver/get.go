package httpserver

import (
	"encoding/json"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	mtx.Lock()
	defer mtx.Unlock()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)

}
