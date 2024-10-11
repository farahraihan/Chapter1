package repository

import (
	"chapter1/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string
	Email    string
	Password string
	Image    string
	IsAdmin  bool
}

func (u *User) ToUserEntity() users.User {
	return users.User{
		ID:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		Email:     u.Email,
		Password:  u.Password,
		IsAdmin:   u.IsAdmin,
		Image:     u.Image,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserQuery(input users.User) User {
	return User{
		Name:     input.Name,
		Phone:    input.Phone,
		Email:    input.Email,
		Password: input.Password,
		IsAdmin:  input.IsAdmin,
		Image:    input.Image,
	}
}
