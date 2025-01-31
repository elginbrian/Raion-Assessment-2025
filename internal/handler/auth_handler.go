package handler

import (
	"log"
	"raion-assessment/internal/service"
	"raion-assessment/pkg/request"
	"raion-assessment/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

func extractToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) <= len("Bearer ") {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Missing or invalid token")
	}
	return authHeader[len("Bearer "):], nil
}

// @Summary Register a new user
// @Description Create a new account by providing a username, email, and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.UserRegistrationRequest true "User registration details"
// @Success 201 {object} response.RegisterResponse "Successful registration response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.UserRegistrationRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return response.ValidationError(c, "Invalid request format")
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("Validation failed: %v", err)
		return response.ValidationError(c, err.Error())
	}

	if err := h.authService.Register(req.Username, req.Email, req.Password); err != nil {
		log.Printf("Registration failed: %v", err)
		return response.Error(c, err.Error())
	}

	return response.Success(c.Status(fiber.StatusCreated), response.RegisterData{
		Message: "User registered successfully",
	})
}

// @Summary Log in a user
// @Description Authenticate user and receive access and refresh tokens.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.UserLoginRequest true "User login details"
// @Success 200 {object} response.LoginResponse "Successful login response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.UserLoginRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return response.ValidationError(c, "Invalid request format")
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("Validation failed: %v", err)
		return response.ValidationError(c, err.Error())
	}

	accessToken, refreshToken, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		log.Printf("Login failed: %v", err)
		return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
	}

	return response.Success(c, response.LoginData{
		AccessToken:  "Bearer " + accessToken,
		RefreshToken: refreshToken,
	})
}

// @Summary Refresh access token
// @Description Obtain a new access token using a valid refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} response.RefreshTokenResponse "New access token"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req request.RefreshTokenRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return response.ValidationError(c, "Invalid request format")
	}

	newAccessToken, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		log.Printf("Token refresh failed: %v", err)
		return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
	}

	return response.Success(c, response.RefreshTokenData{
		AccessToken: "Bearer " + newAccessToken,
	})
}

// @Summary Get current user info
// @Description Retrieve logged-in user's details using an access token.
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.GetCurrentUserResponse "User details"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /auth/current-user [get]
func (h *AuthHandler) GetUserInfo(c *fiber.Ctx) error {
	token, err := extractToken(c)
	if err != nil {
		return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
	}

	ctx := c.Context()
	user, err := h.authService.GetCurrentUser(ctx, token)
	if err != nil {
		log.Printf("Error fetching user info: %v", err)
		return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
	}

	return response.Success(c, response.User{
		ID:        user.ID,
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// @Summary Change your password
// @Description Update your password securely. You need to be logged in and provide your old password along with the new one. Include your JWT token in the Authorization header.
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body request.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} response.ChangePasswordData "Password changed successfully"
// @Failure 400 {object} response.ErrorResponse "Validation error or invalid request format"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /auth/change-password [put]
func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	token, err := extractToken(c)
	if err != nil {
		return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
	}

	ctx := c.Context()
	user, err := h.authService.GetCurrentUser(ctx, token)
	if err != nil {
		log.Printf("Error fetching user info for password change: %v", err)
		return response.Error(c.Status(fiber.StatusUnauthorized), "Unauthorized: "+err.Error())
	}

	var req request.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return response.ValidationError(c, "Invalid request format")
	}

	if err := h.validate.Struct(req); err != nil {
		log.Printf("Validation failed: %v", err)
		return response.ValidationError(c, err.Error())
	}

	if err := h.authService.ChangePassword(user.ID, req.OldPassword, req.NewPassword); err != nil {
		log.Printf("Password change failed: %v", err)
		return response.Error(c, err.Error())
	}

	return response.Success(c, response.ChangePasswordData{
		Message: "Password changed successfully",
	})
}