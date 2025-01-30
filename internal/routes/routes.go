package routes

import (
	"raion-battlepass/config"
	"raion-battlepass/internal/handler"
	"raion-battlepass/internal/middleware"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
	authHandler *handler.AuthHandler,
	postHandler *handler.PostHandler,
	commentHandler *handler.CommentHandler,
	likeHandler *handler.LikeHandler, 
	jwtSecret string,
) {
	config.InitMetrics()
	app.Use(config.PrometheusMiddleware)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	app.Static("/uploads", "./public/uploads")

	app.Get("/", redirectToDocs)
	app.Get("/api", redirectToDocs)
	app.Get("/docs", redirectToDocs)
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	setupUserRoutes(app, userHandler)
	setupAuthRoutes(app, authHandler, jwtSecret)
	setupPostRoutes(app, postHandler)
	setupSearchRoutes(app, userHandler, postHandler)
	setupCommentRoutes(app, commentHandler)
	setupLikeRoutes(app, likeHandler)  

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code": fiber.StatusNotFound,  
			"status": "error",
			"message": "The route you requested does not exist. Please check the URL and try again.",
		})
	})
}

func redirectToDocs(c *fiber.Ctx) error {
	return c.Redirect("/docs/index.html")
}

func setupLikeRoutes(app *fiber.App, handler *handler.LikeHandler) {
	likeGroup := app.Group("/api/v1/posts")
	likeGroup.Post("/:post_id/like", handler.LikePost)    
	likeGroup.Post("/:post_id/unlike", handler.UnlikePost) 
}

func setupSearchRoutes(app *fiber.App, userHandler *handler.UserHandler, postHandler *handler.PostHandler) {
	searchGroup := app.Group("/api/v1/search")
	searchGroup.Get("/users", userHandler.SearchUsers)
	searchGroup.Get("/posts", postHandler.SearchPosts)
}

func setupUserRoutes(app *fiber.App, handler *handler.UserHandler) {
	userGroup := app.Group("/api/v1/users")
	userGroup.Put("/", handler.UpdateUser)
	userGroup.Get("/", handler.GetAllUsers)
	userGroup.Get("/:id", handler.GetUserByID)
}

func setupAuthRoutes(app *fiber.App, handler *handler.AuthHandler, jwtSecret string) {
	authGroup := app.Group("/api/v1/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
	authGroup.Get("/current-user", handler.GetUserInfo, middleware.TokenValidationMiddleware(jwtSecret))
	authGroup.Put("/change-password", handler.ChangePassword, middleware.TokenValidationMiddleware(jwtSecret))
}

func setupPostRoutes(app *fiber.App, handler *handler.PostHandler) {
	postGroup := app.Group("/api/v1/posts")
	postGroup.Post("/", handler.CreatePost)
	postGroup.Put("/:id", handler.UpdatePost)
	postGroup.Delete("/:id", handler.DeletePost)
	postGroup.Get("/", handler.GetAllPosts)
	postGroup.Get("/:id", handler.GetPostByID)
	postGroup.Get("/user/:user_id", handler.GetPostsByUserID)
}

func setupCommentRoutes(app *fiber.App, handler *handler.CommentHandler) {
	commentGroup := app.Group("/api/v1/posts/:post_id/comments")
	commentGroup.Get("/", handler.GetCommentsByPostID)     
	commentGroup.Post("/", handler.CreateComment)         
	commentGroup.Delete("/:id", handler.DeleteComment)    
}