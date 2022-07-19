package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:manish@localhost:5432/sls?sslmode=disable"
)

var (
	DB *sql.DB
)

type Link struct {
	Id        int    `json:"id"`
	ShortLink string `json:"shortlink"`
	LongLink  string `json:"longlink"`
}

func setupRoutes(r *gin.Engine) {
	r.POST("short_link/create", createHandler)
	r.GET("/:id", redirectHandler)
}

func main() {
	createDBConnection()
	r := gin.Default()
	setupRoutes(r)
	r.Run() //

}

func createDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	} else {
		fmt.Println("Connected to database")

	}
	// defer DB.Close()
}
