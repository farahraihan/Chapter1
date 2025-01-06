package handler

import "chapter1/internal/features/quotes"

type AddQuoteRequest struct {
	Content  string `json:"content" form:"content"`
	Caption  string `json:"caption" form:"caption"`
	MemberID uint   `json:"memberID" form:"memberID"`
	BookID   uint   `json:"bookID" form:"bookID"`
}

func ToQuoteModel(qr AddQuoteRequest) quotes.Quote {
	return quotes.Quote{
		Content: qr.Content,
		Caption: qr.Caption,
		BookID:  qr.BookID,
	}
}
