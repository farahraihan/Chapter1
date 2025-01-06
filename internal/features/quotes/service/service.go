package service

import (
	"chapter1/internal/features/quotes"
	"chapter1/internal/features/users"
	"errors"
	"log"
)

type QuoteServices struct {
	qry      quotes.QQuery
	uService users.UService
}

func NewQuoteServices(q quotes.QQuery, u users.UService) quotes.QService {
	return &QuoteServices{
		qry:      q,
		uService: u,
	}
}

func (qs *QuoteServices) AddQuote(newQuote quotes.Quote) error {
	err := qs.qry.AddQuote(newQuote)
	if err != nil {
		log.Println("add quote query error: ", err)
		return errors.New("failed to add quote, please try again later")
	}

	err = qs.uService.AddPoints(newQuote.MemberID, 100)
	if err != nil {
		log.Println("add user point query error: ", err)
		return errors.New("failed to add user point, please try again later")
	}

	return nil
}

func (qs *QuoteServices) GetAllQuotes(limit uint, page uint) ([]quotes.Quote, uint, error) {
	quotes, totalItems, err := qs.qry.GetAllQuotes(limit, page)

	if err != nil {
		log.Println("get all quotes query error: ", err)
		return nil, 0, errors.New("failed to retrieve quotes, please try again later")
	}

	return quotes, totalItems, nil
}
