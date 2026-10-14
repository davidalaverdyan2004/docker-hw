package httpserver

import (
	"employees/db"
	"errors"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	err = db.DeleteEmployee(r.Context(), h.pool, id)
	if errors.Is(err, db.ErrNotFound) {
		http.Error(w, "работник не найден", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
