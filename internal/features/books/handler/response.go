package handler

import (
	"chapter1/internal/features/books"
	"time"
)

type BookResponse struct {
	ID              uint      `json:"bookID"`
	Title           string    `json:"title"`
	Category        string    `json:"category"`
	Author          string    `json:"author"`
	Stock           uint      `json:"stock"`
	Price           float64   `json:"price"`
	Blurb           string    `json:"blurb"`
	Rating          string    `json:"rating"`
	Publisher       string    `json:"publisher"`
	PublicationYear string    `json:"publicationYear"`
	CoverImage      string    `json:"coverImage"`
	AdminID         uint      `json:"adminID"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func ToBookResponse(input books.Book) BookResponse {
	return BookResponse{
		ID:              input.ID,
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
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
	}
}

func ToBookResponses(books []books.Book) []BookResponse {
	responses := make([]BookResponse, len(books))
	for i, book := range books {
		responses[i] = ToBookResponse(book)
	}
	return responses
}
