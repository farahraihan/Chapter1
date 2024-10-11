package users

import (
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        uint
	Name      string
	Phone     string
	Email     string
	Password  string
	Image     string
	IsAdmin   bool
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

type UQuery interface {
	Login(email string) (User, error)
	Register(newAdmins User) error
	UpdateUser(userID uint, updatedUser User) error
	DeleteUser(userID uint) error
	GetUserByID(userID uint) (User, error)
	GetAllUsers(limit int, page int, search string) ([]User, int, error)
	IsAdmin(userID uint) (bool, error)
}

type UService interface {
	Login(email string, password string) (User, string, error)
	Register(newAdmins User, src multipart.File, filename string) error
	UpdateUser(userID uint, updatedUser User, src multipart.File, filename string) error
	DeleteUser(userID uint, memberID uint) error
	GetUserByID(userID uint) (User, error)
	GetAllUsers(userID uint, limit int, page int, search string) ([]User, int, error)
}

type UHandler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
}
