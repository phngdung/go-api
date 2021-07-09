package main

import (
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main.go/database"
	"main.go/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	database.Connect()
	r := gin.Default()
	routes.Setup(r)
	r.Run() // listen and serve on localhost:8080
}
