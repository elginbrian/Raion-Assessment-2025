package main

import (
	"raion-assessment/config"
	"raion-assessment/internal/di"
	"raion-assessment/internal/routes"

	_ "raion-assessment/docs"
)

// @title RAION ASSESSMENT API
// @version 1.0.0
// @description This is a RESTful API for a simple social media application.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8084/api/v1
// @BasePath /
func main() {
	serverPort := config.GetServerPort()
	jwtSecret := config.GetJWTSecret()
	refreshSecret := config.GetRefreshSecret()

	db := config.InitDatabase()
	defer db.Close()

	container := di.NewContainer(db, jwtSecret, refreshSecret)
	app := config.SetupFiber()
	routes.SetupRoutes(app, *container, jwtSecret)

	config.StartServer(app, serverPort)
}