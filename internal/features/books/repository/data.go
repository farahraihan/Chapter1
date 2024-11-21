package repository

import (
	"chapter1/internal/features/books"
	"chapter1/internal/features/users"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
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
}

func (b *Book) ToBookEntity() books.Book {
	return books.Book{
		ID:              b.ID,
		Title:           b.Title,
		Category:        b.Category,
		Author:          b.Author,
		Stock:           b.Stock,
		Price:           b.Price,
		Blurb:           b.Blurb,
		Rating:          b.Rating,
		Publisher:       b.Publisher,
		PublicationYear: b.PublicationYear,
		CoverImage:      b.CoverImage,
		AdminID:         b.AdminID,
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
	}
}

func ToBookQuery(input books.Book) Book {
	return Book{
		Title:           input.Title,
		Category:        input.Category,
		Author:          input.Author,
		Stock:           input.Stock,
		Price:           input.Price,
		Blurb:           input.Blurb,
		Rating:          input.Rating,
		Publisher:       input.Publisher,
		PublicationYear: input.PublicationYear,
		CoverImage:      input.CoverImage,
		AdminID:         input.AdminID,
	}
}
