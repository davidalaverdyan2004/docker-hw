package httpserver

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Position string `json:"position"`
}
