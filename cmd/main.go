package main

import (
	"context"
	"log"
	"raion-assessment/config"
	"raion-assessment/internal/di"
	"raion-assessment/internal/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/jackc/pgx/v4/pgxpool"

	_ "raion-assessment/docs"
)

// @title RAION ASSESSMENT API
// @version 1.0.0
// @description This is a RESTful API for a simple social media application. It allows users to manage their posts, including creating, updating, and deleting posts, and provides authentication using JWT. The API is built using the Fiber framework and interacts with a PostgreSQL database.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:8084/api/v1
// @BasePath /
func main() {
	serverPort := config.GetServerPort()
	databaseURL := config.GetDatabaseURL()
	jwtSecret := config.GetJWTSecret()
	refreshSecret := config.GetRefreshSecret()

	db, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Running database migrations...")
	if err := config.RunSQLMigrations(db); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	container := di.NewContainer(db, jwtSecret, refreshSecret)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRoutes(app, container.UserHandler, container.AuthHandler, container.PostHandler, container.CommentHandler, container.LikeHandler, jwtSecret)

	log.Printf("Server is running on port %s", serverPort)
	if err := app.Listen(serverPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}