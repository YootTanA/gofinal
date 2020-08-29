package task

import "database/sql"

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type Handler struct {
	DB *sql.DB
}
