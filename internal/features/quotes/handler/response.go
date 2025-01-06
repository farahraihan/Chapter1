package handler

import (
	"chapter1/internal/features/quotes"
	"time"
)

type QuoteResponse struct {
	ID         uint      `json:"quoteID"`
	Content    string    `json:"content"`
	Caption    string    `json:"caption"`
	MemberID   uint      `json:"memberID"`
	MemberName string    `json:"memberName"`
	BookID     uint      `json:"bookID"`
	BookTitle  string    `json:"bookTitle"`
	BookAuthor string    `json:"bookAuthor"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func ToQuoteResponse(input quotes.Quote) QuoteResponse {
	return QuoteResponse{
		ID:         input.ID,
		Content:    input.Content,
		Caption:    input.Caption,
		MemberID:   input.MemberID,
		MemberName: input.MemberName,
		BookID:     input.BookID,
		BookTitle:  input.BookTitle,
		BookAuthor: input.BookAuthor,
	}
}

func ToQuoteResponses(quotes []quotes.Quote) []QuoteResponse {
	responses := make([]QuoteResponse, len(quotes))
	for i, quote := range quotes {
		responses[i] = ToQuoteResponse(quote)
	}
	return responses
}
