package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"secure-web-app/backend/routes"
	"secure-web-app/backend/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	utils.InitializeDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	defer utils.CloseDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}