package handler

import (
	"chapter1/internal/features/users"
	"time"
)

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	IsAdmin   bool      `json:"is_admin"`
	Image     string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToLoginResponse(result users.User, token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}

func ToUserResponse(input users.User) UserResponse {
	return UserResponse{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Phone:     input.Phone,
		IsAdmin:   input.IsAdmin,
		Image:     input.Image,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func ToUserResponses(users []users.User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = ToUserResponse(user)
	}
	return responses
}
