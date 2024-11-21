package books

import (
	"chapter1/internal/features/users"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Book struct {
	ID              uint
	Title           string
	Category        string
	Author          string
	Stock           uint
	Price           float64
	Blurb           string
	Rating          string
	Publisher       string
	PublicationYear string
	CoverImage      string
	AdminID         uint
	User            users.User `gorm:"foreignKey:AdminID"`
	CreatedAt       time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time  `gorm:"default:current_timestamp"`
}

type BQuery interface {
	AddBook(newBooks Book) error
	UpdateBook(bookID uint, updatedBook Book) error
	DeleteBook(bookID uint) error
	GetBookByID(bookID uint) (Book, error)
	GetAllBooks(limit uint, page uint, search string) ([]Book, uint, error)
}

type BService interface {
	AddBook(userID uint, newBooks Book, src multipart.File, filename string) error
	UpdateBook(userID uint, bookID uint, updatedBook Book, src multipart.File, filename string) error
	DeleteBook(userID uint, bookID uint) error
	GetBookByID(bookID uint) (Book, error)
	GetAllBooks(limit uint, page uint, search string) ([]Book, uint, error)
}

type BHandler interface {
	AddBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
	GetBookByID() echo.HandlerFunc
	GetAllBooks() echo.HandlerFunc
}
