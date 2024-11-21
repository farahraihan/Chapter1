package repository

import (
	"chapter1/internal/features/books"

	"gorm.io/gorm"
)

type BookQuery struct {
	db *gorm.DB
}

func NewBookQuery(connect *gorm.DB) books.BQuery {
	return &BookQuery{
		db: connect,
	}
}

func (bq *BookQuery) AddBook(newBook books.Book) error {
	cnvData := ToBookQuery(newBook)
	qry := bq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (bq *BookQuery) UpdateBook(bookID uint, updatedBook books.Book) error {
	cnvData := ToBookQuery(updatedBook)
	qry := bq.db.Where("id = ?", bookID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (bq *BookQuery) DeleteBook(bookID uint) error {
	qry := bq.db.Where("id = ?", bookID).Delete(&Book{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (bq *BookQuery) GetBookByID(bookID uint) (books.Book, error) {
	var book Book

	qry := bq.db.Preload("User").First(&book, bookID)

	if qry.Error != nil {
		return books.Book{}, qry.Error
	}

	return book.ToBookEntity(), nil
}

func (bq *BookQuery) GetAllBooks(limit uint, page uint, search string) ([]books.Book, uint, error) {
	var booksList []Book
	var totalItem int64

	offset := (page - 1) * limit

	qry := bq.db.Model(&Book{}).Where("title LIKE ?", "%"+search+"%").Count(&totalItem)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	qry = bq.db.Where("title LIKE ?", "%"+search+"%").Limit(int(limit)).Offset(int(offset)).Find(&booksList)
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	booksEntities := make([]books.Book, len(booksList))
	for i, book := range booksList {
		booksEntities[i] = book.ToBookEntity()
	}

	return booksEntities, uint(totalItem), nil
}
