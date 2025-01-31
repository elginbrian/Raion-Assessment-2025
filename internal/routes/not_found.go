package routes

import "github.com/gofiber/fiber/v2"

func setupNotFoundHandler(app *fiber.App) {
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"status":  "error",
			"message": "The route you requested does not exist. Please check the URL and try again.",
		})
	})
}