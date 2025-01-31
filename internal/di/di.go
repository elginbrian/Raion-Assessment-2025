package di

import (
	"raion-assessment/internal/handler"
	"raion-assessment/internal/repository"
	"raion-assessment/internal/service"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Container struct {
	UserHandler    *handler.UserHandler
	AuthHandler    *handler.AuthHandler
	PostHandler    *handler.PostHandler
	CommentHandler *handler.CommentHandler
	LikeHandler    *handler.LikeHandler
}

func NewContainer(db *pgxpool.Pool, jwtSecret string, refreshSecret string) *Container {
	// Repositories
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db) 
	likeRepo := repository.NewLikeRepository(db) 

	// Services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, authRepo, jwtSecret, refreshSecret)
	postService := service.NewPostService(postRepo) 
	commentService := service.NewCommentService(commentRepo)
	likeService := service.NewLikeService(likeRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService, authService)
	authHandler := handler.NewAuthHandler(authService)
	postHandler := handler.NewPostHandler(postService, authService) 
	commentHandler := handler.NewCommentHandler(commentService, authService)
	likeHandler := handler.NewLikeHandler(likeService, authService)

	return &Container{
		UserHandler: userHandler,
		AuthHandler: authHandler,
		PostHandler: postHandler, 
		CommentHandler: commentHandler,
		LikeHandler: likeHandler,
	}
}
