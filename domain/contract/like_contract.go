package domain

import (
	"context"
	domain "raion-assessment/domain/entity"
)

type ILikeRepository interface {
	GetLikesByPostID(ctx context.Context, postID string) ([]domain.Like, error)
	GetLikesByUserID(ctx context.Context, userID string) ([]domain.Like, error)
	AddLike(ctx context.Context, like domain.Like) (*domain.Like, error)
	RemoveLike(ctx context.Context, userID, postID string) error
}

type ILikeService interface {
	GetLikesByPostID(postID string) ([]domain.Like, error)
	GetLikesByUserID(userID string) ([]domain.Like, error)
	AddLike(userID, postID string) (domain.Like, error)
	RemoveLike(userID, postID string) error
}