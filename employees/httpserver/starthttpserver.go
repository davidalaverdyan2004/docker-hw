package httpserver

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StartHttpServer(pool *pgxpool.Pool) error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Docker"))
	})

	h := &Handler{pool: pool}

	http.HandleFunc("GET /employees", h.GetEmployees)

	http.HandleFunc("POST /employees", h.CreateEmployee)

	http.HandleFunc("DELETE /employees/{id}", DeleteEmployee)

	err := http.ListenAndServe(":5050", nil)
	if err == http.ErrServerClosed {
		return nil
	} else {
		return err
	}
}
