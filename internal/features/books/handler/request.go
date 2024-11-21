package handler

import "chapter1/internal/features/books"

type AddBookRequest struct {
	Title           string  `json:"title" form:"title"`
	Category        string  `json:"category" form:"category"`
	Author          string  `json:"author" form:"author"`
	Stock           uint    `json:"stock" form:"stock"`
	Price           float64 `json:"price" form:"price"`
	Blurb           string  `json:"blurb" form:"blurb"`
	Rating          string  `json:"rating" form:"rating"`
	Publisher       string  `json:"publisher" form:"publisher"`
	PublicationYear string  `json:"publicationYear" form:"publicationYear"`
	CoverImage      string  `json:"coverImage" form:"coverImage"`
	AdminID         uint    `json:"adminID" form:"adminID"`
}

func ToBookModel(br AddBookRequest) books.Book {
	return books.Book{
		Title:           br.Title,
		Category:        br.Category,
		Author:          br.Author,
		Stock:           br.Stock,
		Price:           br.Price,
		Blurb:           br.Blurb,
		Rating:          br.Rating,
		Publisher:       br.Publisher,
		PublicationYear: br.PublicationYear,
		CoverImage:      br.CoverImage,
		AdminID:         br.AdminID,
	}
}
