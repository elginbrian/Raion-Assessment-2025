package service

import (
	"context"
	"raion-battlepass/internal/domain"
	"raion-battlepass/internal/repository"
)

type LikeService interface {
	GetLikesByPostID(postID string) ([]domain.Like, error)
	GetLikesByUserID(userID string) ([]domain.Like, error)
	AddLike(userID, postID string) (domain.Like, error)
	RemoveLike(userID, postID string) error
}

type likeService struct {
	likeRepo repository.LikeRepository
}

func NewLikeService(repo repository.LikeRepository) LikeService {
	return &likeService{likeRepo: repo}
}

func (s *likeService) GetLikesByPostID(postID string) ([]domain.Like, error) {
	ctx := context.Background()
	likes, err := s.likeRepo.GetLikesByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if len(likes) == 0 {
		return nil, domain.ErrNotFound
	}
	return likes, nil
}

func (s *likeService) GetLikesByUserID(userID string) ([]domain.Like, error) {
	ctx := context.Background()
	likes, err := s.likeRepo.GetLikesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(likes) == 0 {
		return nil, domain.ErrNotFound
	}
	return likes, nil
}

func (s *likeService) AddLike(userID, postID string) (domain.Like, error) {
	ctx := context.Background()
	like := domain.Like{
		UserID: userID,
		PostID: postID,
	}
	createdLike, err := s.likeRepo.AddLike(ctx, like)
	if err != nil {
		return domain.Like{}, err
	}
	return *createdLike, nil
}

func (s *likeService) RemoveLike(userID, postID string) error {
	ctx := context.Background()
	err := s.likeRepo.RemoveLike(ctx, userID, postID)
	if err != nil {
		return err
	}
	return nil
}