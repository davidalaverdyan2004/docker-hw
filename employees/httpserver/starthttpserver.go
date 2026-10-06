package httpserver

import (
	"net/http"
)

func StartHttpServer() error {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Docker"))
	})

	http.HandleFunc("GET /employees", GetEmployees)

	http.HandleFunc("POST /employees", CreateEmployee)

	http.HandleFunc("DELETE /employees/{id}", DeleteEmployee)

	err := http.ListenAndServe(":5050", nil)
	if err == http.ErrServerClosed {
		return nil
	} else {
		return err
	}
}
