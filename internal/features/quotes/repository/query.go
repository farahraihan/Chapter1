package repository

import (
	"chapter1/internal/features/quotes"

	"gorm.io/gorm"
)

type QuoteQuery struct {
	db *gorm.DB
}

func NewQuoteQuery(connect *gorm.DB) quotes.QQuery {
	return &QuoteQuery{
		db: connect,
	}
}

func (qq *QuoteQuery) AddQuote(newQuote quotes.Quote) error {
	cnvData := ToQuoteQuery(newQuote)
	qry := qq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (qq *QuoteQuery) GetAllQuotes(limit uint, page uint) ([]quotes.Quote, uint, error) {
	var quotesList []Quote
	var totalItem int64

	offset := (page - 1) * limit

	qry := qq.db.Model(&Quote{}).Count(&totalItem)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	qry = qq.db.Preload("User").Preload("Book").Limit(int(limit)).Offset(int(offset)).Find(&quotesList)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	quotesEntities := make([]quotes.Quote, len(quotesList))
	for i, quote := range quotesList {
		quotesEntities[i] = quote.ToQuoteEntity()
	}

	return quotesEntities, uint(totalItem), nil
}
