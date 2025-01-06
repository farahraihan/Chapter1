package handler

import (
	"chapter1/internal/features/quotes"
	"chapter1/internal/helpers"
	"chapter1/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type QuoteHandler struct {
	srv quotes.QService
	tu  utils.JwtUtilityInterface
}

func NewQuoteHandler(s quotes.QService, t utils.JwtUtilityInterface) quotes.QHandler {
	return &QuoteHandler{
		srv: s,
		tu:  t,
	}
}

func (qh *QuoteHandler) AddQuote() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := qh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
		}

		var input AddQuoteRequest

		err := c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add quote request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newQuote := ToQuoteModel(input)
		newQuote.MemberID = uint(userID)

		err = qh.srv.AddQuote(newQuote)
		if err != nil {
			log.Print("failed add a quote")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed add a quote", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "success add a quote", nil))
	}
}

// lanjut function kedua
func (qh *QuoteHandler) GetAllQuotes() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 0 {
			page = 1
		}

		quotes, totalItems, err := qh.srv.GetAllQuotes(uint(limit), uint(page))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve quote", nil))
		}

		responseData := ToQuoteResponses(quotes)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all books", responseData, meta))
	}
}
