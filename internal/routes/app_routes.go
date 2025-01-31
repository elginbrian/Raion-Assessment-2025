package routes

import (
	"raion-assessment/internal/di"

	_ "raion-assessment/docs"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, container di.Container, jwtSecret string) {
	setupStaticRoutes(app)
	setupDocsRoutes(app)
	setupAPIRoutes(app, container, jwtSecret)
	setupNotFoundHandler(app)
}