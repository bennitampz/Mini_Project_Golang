package main

import (
	"fmt"
	"log"
	"os"

	"tokoku_go/config"
	"tokoku_go/database/migrations"
	"tokoku_go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	migrations.Migrate()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
