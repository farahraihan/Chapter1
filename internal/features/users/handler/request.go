package handler

import "chapter1/internal/features/users"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Image    string `json:"image" form:"image"`
}

func RegisterToUser(ur RegisterRequest) users.User {
	return users.User{
		Name:     ur.Name,
		Phone:    ur.Phone,
		Email:    ur.Email,
		Password: ur.Password,
		Image:    ur.Image,
	}
}
