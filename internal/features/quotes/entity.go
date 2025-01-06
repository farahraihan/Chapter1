package quotes

import (
	"chapter1/internal/features/books"
	"chapter1/internal/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Quote struct {
	ID         uint
	Content    string
	Caption    string
	BookID     uint
	MemberID   uint
	MemberName string
	BookTitle  string
	BookAuthor string
	User       users.User `gorm:"foreignKey:MemberID"`
	Book       books.Book `gorm:"foreignKey:BookID"`
	CreatedAt  time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time  `gorm:"default:current_timestamp"`
}

type QQuery interface {
	AddQuote(newQuote Quote) error
	GetAllQuotes(limit uint, page uint) ([]Quote, uint, error)
}

type QService interface {
	AddQuote(newQuote Quote) error
	GetAllQuotes(limit uint, page uint) ([]Quote, uint, error)
}

type QHandler interface {
	AddQuote() echo.HandlerFunc
	GetAllQuotes() echo.HandlerFunc
}
