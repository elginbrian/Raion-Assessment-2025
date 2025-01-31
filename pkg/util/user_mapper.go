package util

import (
	entity "raion-assessment/domain/entity"
	"raion-assessment/pkg/response"
)

func MapToUserResponse(user entity.User) response.User {
	return response.User{
		ID:        user.ID,
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}