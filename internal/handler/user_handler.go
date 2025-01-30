package handler

import (
	"fmt"
	"log"
	"raion-assessment/internal/domain"
	"raion-assessment/internal/service"
	"raion-assessment/pkg/request"
	"raion-assessment/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(userService service.UserService, authService service.AuthService) *UserHandler {
    return &UserHandler{
        userService: userService,
        authService: authService,
    }
}

func mapToUserResponse(user domain.User) domain.UserResponse {
	return domain.UserResponse{
		ID:        user.ID,
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users from the database.
// @Tags users
// @Produce json
// @Success 200 {object} response.GetAllUsersResponse "Successful fetch users response"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.FetchAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return response.Error(c, "Error fetching users", fiber.StatusInternalServerError)
	}

	var userResponses []domain.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, mapToUserResponse(user))
	}

	return response.Success(c, userResponses)
}

// GetUserByID godoc
// @Summary Get user details by ID
// @Description Retrieve the details of a specific user by their ID.
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.GetUserByIDResponse "Successful fetch user by ID response"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.FetchUserByID(id)
	if err != nil {
		log.Printf("Error fetching user by ID: %v", err)
		return response.Error(c, "User not found", fiber.StatusNotFound)
	}

	return response.Success(c, mapToUserResponse(user))
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update the username of the authenticated user.
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.UpdateUserRequest true "Updated username"
// @Security BearerAuth
// @Success 200 {object} response.UpdateUserResponse "Successful update user response"
// @Failure 400 {object} response.ErrorResponse "Validation error"
// @Failure 401 {object} response.ErrorResponse "Unauthorized or invalid token"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /users [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	token, err := extractToken(c)
	if err != nil {
		return response.Error(c.Status(fiber.StatusUnauthorized), "Unauthorized: "+err.Error())
	}

	ctx := c.Context()
	user, err := h.authService.GetCurrentUser(ctx, token)
	if err != nil {
		log.Printf("Unauthorized access: %v", err)
		return response.Error(c.Status(fiber.StatusUnauthorized), "Unauthorized: "+err.Error())
	}

	var payload request.UpdateUserRequest
	if err := c.BodyParser(&payload); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return response.ValidationError(c, "Invalid input, expected JSON with 'username'")
	}

	if len(payload.Username) < 3 || len(payload.Username) > 50 {
		return response.ValidationError(c, "Username must be between 3 and 50 characters")
	}

	updatedUser := domain.User{
		ID:        user.ID,
		Name:      payload.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	updatedUser, err = h.userService.UpdateUser(user.ID, updatedUser)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return response.Error(c.Status(fiber.StatusInternalServerError), fmt.Sprintf("Error updating user: %v", err))
	}

	return response.Success(c, mapToUserResponse(updatedUser))
}

// SearchUsers godoc
// @Summary Search users
// @Description Search for users by their name or email.
// @Tags search
// @Produce json
// @Param query query string true "Search query"
// @Success 200 {array} domain.UserResponse "Successful search response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /search/users [get]
func (h *UserHandler) SearchUsers(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return response.Error(c, "Query parameter is required", fiber.StatusBadRequest)
	}

	log.Printf("Received search query: %s", query)

	users, err := h.userService.SearchUsers(query)
	if err != nil {
		if err.Error() == "no users found" {
			return response.Error(c, "No users found for the given search query", fiber.StatusNotFound)
		}
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	var userResponses []domain.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, mapToUserResponse(user))
	}

	return response.Success(c, userResponses)
}