package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	db := os.Getenv("DB")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, db)

	// API request and response
	router := gin.Default()

	// enable cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {

		// MYSQL DATABASE CONNECTIONS
		db, err := sql.Open("mysql", dbURI)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		// RANDOM NUMBERS GENERATES
		var num int = rand.Intn(10) + 1

		// MYSQL QUERY
		var content string
		err = db.QueryRow("SELECT Content FROM quote WHERE ID = ?", num).Scan(&content)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(200, gin.H{
			"message": content,
		})

	})

	router.Run(":8080")
}
