package handler

import (
	"chapter1/internal/features/books"
	"chapter1/internal/helpers"
	"chapter1/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	srv books.BService
	tu  utils.JwtUtilityInterface
}

func NewBookHandler(s books.BService, t utils.JwtUtilityInterface) books.BHandler {
	return &BookHandler{
		srv: s,
		tu:  t,
	}
}

func (bh *BookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := bh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		image, err := c.FormFile("coverImage")
		if err != nil {
			log.Println("failed to get image file")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid image file", nil))
		}

		src, err := image.Open()
		if err != nil {
			log.Println("failed to open image file")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process the image", nil))
		}
		defer src.Close()

		var input AddBookRequest

		err = c.Bind(&input)
		if err != nil {
			log.Print("failed to bind add book request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newBook := ToBookModel(input)
		newBook.AdminID = uint(userID)

		err = bh.srv.AddBook(uint(userID), newBook, src, image.Filename)
		if err != nil {
			log.Print("failed add a book")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed add a book", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "success add a book", nil))

	}

}

func (bh *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := bh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		image, err := c.FormFile("coverImage")
		if err != nil {
			log.Println("failed to get image file")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid image file", nil))
		}

		src, err := image.Open()
		if err != nil {
			log.Println("failed to open image file")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process the image", nil))
		}
		defer src.Close()

		bookIDStr := c.Param("id")
		bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
		if err != nil {
			log.Println("invalid book ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid book ID", nil))
		}

		var input AddBookRequest

		err = c.Bind(&input)
		if err != nil {
			log.Print("failed to bind add book request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		updatedBook := ToBookModel(input)
		updatedBook.AdminID = uint(userID)

		err = bh.srv.UpdateBook(uint(userID), uint(bookID), updatedBook, src, image.Filename)
		if err != nil {
			log.Println("failed to update book")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to update", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "book was successfully updated", nil))

	}
}

func (bh *BookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := bh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		bookIDStr := c.Param("id")
		bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
		if err != nil {
			log.Println("invalid book ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid book ID", nil))
		}

		err = bh.srv.DeleteBook(uint(userID), uint(bookID))
		if err != nil {
			log.Println("failed to delete book")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "book was successfully deleted", nil))

	}
}

func (bh *BookHandler) GetBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookIDStr := c.Param("id")
		bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
		if err != nil {
			log.Println("invalid book ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid book ID", nil))
		}

		book, err := bh.srv.GetBookByID(uint(bookID))
		if err != nil {
			log.Println("failed to get book by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve book", nil))
		}

		BookResponse := ToBookResponse(book)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "book was successfully retrieved", BookResponse))

	}
}

func (bh *BookHandler) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10 // default limit
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 0 {
			page = 1 // default page
		}

		search := c.QueryParam("search")
		if search == "" {
			search = "" // default search string
		}

		books, totalItems, err := bh.srv.GetAllBooks(uint(limit), uint(page), search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve books", nil))
		}

		responseData := ToBookResponses(books)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all books", responseData, meta))

	}
}
