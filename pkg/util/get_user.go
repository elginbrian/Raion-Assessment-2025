package util

import (
	contract "raion-assessment/domain/contract"
	entity "raion-assessment/domain/entity"
	"raion-assessment/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func GetUserFromToken(c *fiber.Ctx, authService contract.IAuthService) (*entity.User, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) <= len("Bearer ") {
		return nil, response.Error(c, "Missing or invalid token", fiber.StatusUnauthorized)
	}
	token := authHeader[len("Bearer "):]
	ctx := c.Context()
	return authService.GetCurrentUser(ctx, token)
}
