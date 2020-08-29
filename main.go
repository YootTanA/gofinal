package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/YootTanA/gofinal/middleware"

	"github.com/YootTanA/gofinal/task"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Auth)

	h := task.Handler{
		DB: db,
	}

	r.POST("/customers", h.CreateCustomerHandler)
	r.GET("/customers/:id", h.GetCustomerById)
	r.GET("/customers", h.GetCustomers)
	r.PUT("/customers/:id", h.PutCustomerById)
	r.DELETE("/customers/:id", h.DeleteCustomerById)

	return r
}

func createTable() {
	createTb := `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
		);`

	_, err := db.Exec(createTb)

	if err != nil {
		log.Fatal("Cannot Create Database")
	}

	fmt.Println("Create Table Success\n")
}

func main() {
	createTable()
	r := setupRouter()
	r.Run(":2009")
}
