package response

type LoginResponse struct {
	Status string    `json:"status"`
	Data   LoginData `json:"data"`
	Code   int       `json:"code"`
}

type LoginData struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Status string       `json:"status"`
	Data   RegisterData `json:"data"`
	Code   int          `json:"code"`
}

type RegisterData struct {
	Message string `json:"message"`
}

type GetCurrentUserResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
	Code   int    `json:"code"`
}

type ChangePasswordResponse struct {
	Status string             `json:"status"`
	Data   ChangePasswordData `json:"data"`
	Code   int                `json:"code"`
}

type ChangePasswordData struct {
	Message string `json:"message"`
}