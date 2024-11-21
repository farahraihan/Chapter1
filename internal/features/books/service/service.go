package service

import (
	"chapter1/internal/features/books"
	"chapter1/internal/features/users"
	"chapter1/internal/utils"
	"errors"
	"log"
	"mime/multipart"
)

type BookServices struct {
	qry        books.BQuery
	jwt        utils.JwtUtilityInterface
	cloudinary utils.CloudinaryUtilityInterface
	uService   users.UService
}

func NewBookServices(q books.BQuery, j utils.JwtUtilityInterface, c utils.CloudinaryUtilityInterface, u users.UService) books.BService {
	return &BookServices{
		qry:        q,
		jwt:        j,
		cloudinary: c,
		uService:   u,
	}
}

func (bs *BookServices) AddBook(userID uint, newBooks books.Book, src multipart.File, filename string) error {
	isAdmin, err := bs.uService.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("delete user permission error: ", err)
		return errors.New("access denied")
	}

	imageURL, err := bs.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		log.Println("image upload failed: ", err)
		return errors.New("failed to upload image, please try again later")
	}
	newBooks.CoverImage = imageURL

	err = bs.qry.AddBook(newBooks)
	if err != nil {
		log.Println("add book query error: ", err)
		return errors.New("failed to add a book, please try again later")
	}

	return nil
}

func (bs *BookServices) UpdateBook(userID uint, bookID uint, updatedBook books.Book, src multipart.File, filename string) error {
	isAdmin, err := bs.uService.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("delete user permission error: ", err)
		return errors.New("access denied")
	}

	imageURL, err := bs.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		log.Println("image upload failed: ", err)
		return errors.New("failed to upload image, please try again later")
	}
	updatedBook.CoverImage = imageURL

	err = bs.qry.UpdateBook(bookID, updatedBook)
	if err != nil {
		log.Println("update book query error")
		return errors.New("update failed, please try again later")
	}

	return nil
}

func (bs *BookServices) DeleteBook(userID uint, bookID uint) error {
	isAdmin, err := bs.uService.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("delete user permission error: ", err)
		return errors.New("access denied")
	}

	err = bs.qry.DeleteBook(bookID)
	if err != nil {
		log.Println("delete book query error")
		return errors.New("delete failed, please try again later")
	}

	return nil
}

func (bs *BookServices) GetBookByID(bookID uint) (books.Book, error) {
	book, err := bs.qry.GetBookByID(bookID)
	if err != nil {
		log.Println("get book by ID query error: ", err)
		return books.Book{}, errors.New("failed to retrieve book, please try again later")
	}

	return book, nil
}

func (bs *BookServices) GetAllBooks(limit uint, page uint, search string) ([]books.Book, uint, error) {
	books, totalItems, err := bs.qry.GetAllBooks(limit, page, search)

	if err != nil {
		log.Println("get all books query error: ", err)
		return nil, 0, errors.New("failed to retrieve books, please try again later")
	}

	return books, totalItems, nil

}
