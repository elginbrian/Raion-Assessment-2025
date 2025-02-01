package di

import (
	graph "raion-assessment/internal/handler/graphql"
	rest "raion-assessment/internal/handler/rest"
	"raion-assessment/internal/repository"
	"raion-assessment/internal/service"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Container struct {
	UserHandler    *rest.UserHandler
	AuthHandler    *rest.AuthHandler
	PostHandler    *rest.PostHandler
	CommentHandler *rest.CommentHandler
	LikeHandler    *rest.LikeHandler
	UserResolver   *graph.UserResolver
	PostResolver   *graph.PostResolver
}

func NewContainer(db *pgxpool.Pool, jwtSecret string, refreshSecret string) *Container {
	// Repositories
	userRepo 	:= repository.NewUserRepository(db)
	authRepo 	:= repository.NewAuthRepository(db)
	postRepo 	:= repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db) 
	likeRepo 	:= repository.NewLikeRepository(db) 

	// Services
	userService 	:= service.NewUserService(userRepo)
	authService 	:= service.NewAuthService(userRepo, authRepo, jwtSecret, refreshSecret)
	postService 	:= service.NewPostService(postRepo) 
	commentService 	:= service.NewCommentService(commentRepo)
	likeService 	:= service.NewLikeService(likeRepo)

	// Handlers
	userHandler 	:= rest.NewUserHandler(userService, authService)
	authHandler 	:= rest.NewAuthHandler(authService)
	postHandler 	:= rest.NewPostHandler(postService, authService) 
	commentHandler 	:= rest.NewCommentHandler(commentService, authService)
	likeHandler 	:= rest.NewLikeHandler(likeService, authService)

	// Resolvers
	userResolver 	:= graph.NewUserResolver(userService, authService)
	postResolver 	:= graph.NewPostResolver(postService, authService)

	return &Container{
		UserHandler: userHandler,
		AuthHandler: authHandler,
		PostHandler: postHandler, 
		CommentHandler: commentHandler,
		LikeHandler: likeHandler,
		UserResolver: userResolver,
		PostResolver: postResolver,
	}
}
