package service

import (
	"context"
	"raion-battlepass/internal/domain"
	"raion-battlepass/internal/repository"
)

type CommentService interface {
	GetCommentsByPostID(postID string) ([]domain.Comment, error)
	GetCommentByID(commentID string) (domain.Comment, error)
	CreateComment(comment domain.Comment) (domain.Comment, error)
	DeleteComment(commentID string) error
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{commentRepo: repo}
}

func (s *commentService) GetCommentsByPostID(postID string) ([]domain.Comment, error) {
	ctx := context.Background()
	comments, err := s.commentRepo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return nil, domain.ErrNotFound
	}
	return comments, nil
}

func (s *commentService) GetCommentByID(commentID string) (domain.Comment, error) {
	ctx := context.Background()
	comment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return domain.Comment{}, err
	}
	if comment == nil {
		return domain.Comment{}, domain.ErrNotFound
	}
	return *comment, nil
}

func (s *commentService) CreateComment(comment domain.Comment) (domain.Comment, error) {
	ctx := context.Background()
	createdComment, err := s.commentRepo.CreateComment(ctx, comment)
	if err != nil {
		return domain.Comment{}, err
	}
	if createdComment == nil {
		return domain.Comment{}, domain.ErrNotFound
	}
	return *createdComment, nil
}

func (s *commentService) DeleteComment(commentID string) error {
	ctx := context.Background()
	err := s.commentRepo.DeleteComment(ctx, commentID)
	if err != nil {
		return err
	}
	return nil
}
