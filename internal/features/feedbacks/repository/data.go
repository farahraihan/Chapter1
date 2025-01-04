package repository

import (
	"chapter1/internal/features/books"
	"chapter1/internal/features/feedbacks"
	"chapter1/internal/features/users"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Content  string
	Rating   uint
	MemberID uint
	BookID   uint
	User     users.User `gorm:"foreignKey:MemberID"`
	Book     books.Book `gorm:"foreignKey:BookID"`
}

func (f *Feedback) ToFeedbackEntity() feedbacks.Feedback {
	return feedbacks.Feedback{
		ID:         f.ID,
		Content:    f.Content,
		Rating:     f.Rating,
		MemberID:   f.MemberID,
		BookID:     f.BookID,
		MemberName: f.User.Name,
		BookTitle:  f.Book.Title,
		CreatedAt:  f.CreatedAt,
		UpdatedAt:  f.UpdatedAt,
	}
}

func ToFeedbackQuery(input feedbacks.Feedback) Feedback {
	return Feedback{
		Content:  input.Content,
		Rating:   input.Rating,
		MemberID: input.MemberID,
		BookID:   input.BookID,
	}
}
