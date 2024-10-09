package handler

import "chapter1/internal/features/users"

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	IsAdmin bool   `json:"is_admin"`
	Image   string `json:"image_url"`
}

func ToLoginResponse(result users.User, token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}

// func ToGetUserResponse(input users.User) UserResponse {
// 	return UserResponse{
// 		ID:      input.ID,
// 		Name:    input.Name,
// 		Email:   input.Email,
// 		Phone:   input.Phone,
// 		IsAdmin: input.IsAdmin,
// 		Image:   input.Image,
// 	}
// }
