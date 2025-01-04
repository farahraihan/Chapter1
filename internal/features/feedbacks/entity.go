package feedbacks

import (
	"chapter1/internal/features/books"
	"chapter1/internal/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Feedback struct {
	ID         uint
	Content    string
	Rating     uint
	MemberID   uint
	BookID     uint
	MemberName string
	BookTitle  string
	User       users.User `gorm:"foreignKey:MemberID"`
	Book       books.Book `gorm:"foreignKey:BookID"`
	CreatedAt  time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time  `gorm:"default:current_timestamp"`
}

type FQuery interface {
	AddFeedback(newFeedback Feedback) error
	DeleteFeedback(userID uint, feedbackID uint) error
	GetAllFeedbacks(limit uint, page uint) ([]Feedback, uint, error)
}

type FService interface {
	AddFeedback(newFeedback Feedback) error
	DeleteFeedback(userID uint, feedbackID uint) error
	GetAllFeedbacks(limit uint, page uint) ([]Feedback, uint, error)
}

type FHandler interface {
	AddFeedback() echo.HandlerFunc
	DeleteFeedback() echo.HandlerFunc
	GetAllFeedbacks() echo.HandlerFunc
}
