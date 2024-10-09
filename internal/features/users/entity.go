package users

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	Name     string
	Phone    string
	Email    string
	Password string
	Image    string
	IsAdmin  bool
}

type UQuery interface {
	Login(email string) (User, error)
	Register(newAdmins User) error
}

type UService interface {
	Login(email string, password string) (User, string, error)
	Register(newAdmins User, src multipart.File, filename string) error
}

type UHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
}
