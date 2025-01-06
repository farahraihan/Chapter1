package repository

import (
	"chapter1/internal/features/books"
	"chapter1/internal/features/quotes"
	"chapter1/internal/features/users"

	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	Content  string
	Caption  string
	BookID   uint
	MemberID uint
	User     users.User `gorm:"foreignKey:MemberID"`
	Book     books.Book `gorm:"foreignKey:BookID"`
}

func (q *Quote) ToQuoteEntity() quotes.Quote {
	return quotes.Quote{
		ID:         q.ID,
		Content:    q.Content,
		Caption:    q.Caption,
		MemberID:   q.MemberID,
		BookID:     q.BookID,
		MemberName: q.User.Name,
		BookTitle:  q.Book.Title,
		BookAuthor: q.Book.Author,
		CreatedAt:  q.CreatedAt,
		UpdatedAt:  q.UpdatedAt,
	}
}

func ToQuoteQuery(input quotes.Quote) Quote {
	return Quote{
		Content:  input.Content,
		Caption:  input.Caption,
		MemberID: input.MemberID,
		BookID:   input.BookID,
	}
}
