package handler

import (
	"chapter1/internal/features/feedbacks"
	"chapter1/internal/helpers"
	"chapter1/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	srv feedbacks.FService
	tu  utils.JwtUtilityInterface
}

func NewFeedbackHandler(s feedbacks.FService, t utils.JwtUtilityInterface) feedbacks.FHandler {
	return &FeedbackHandler{
		srv: s,
		tu:  t,
	}
}

func (fh *FeedbackHandler) AddFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := fh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var input AddFeedbackRequest

		err := c.Bind(&input)
		if err != nil {
			log.Println("failed to bind add feedback request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newFeedback := ToFeedbackModel(input)
		newFeedback.MemberID = uint(userID)

		err = fh.srv.AddFeedback(newFeedback)
		if err != nil {
			log.Print("failed add a feedback")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed add a feedback", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "success add a feedback", nil))

	}
}

func (fh *FeedbackHandler) DeleteFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := fh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		feedbackIDStr := c.Param("id")
		feedbackID, err := strconv.ParseUint(feedbackIDStr, 10, 32)
		if err != nil {
			log.Println("invalid book ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid feedback ID", nil))
		}

		err = fh.srv.DeleteFeedback(uint(userID), uint(feedbackID))
		if err != nil {
			log.Println("failed to delete feedback")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to delete", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "feedback was successfully deleted", nil))

	}
}

func (fh *FeedbackHandler) GetAllFeedbacks() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 0 {
			page = 1
		}

		feedbacks, totalItems, err := fh.srv.GetAllFeedbacks(uint(limit), uint(page))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve feedbacks", nil))
		}

		responseData := ToFeedbackResponses(feedbacks)
		meta := helpers.Meta{
			TotalItems:   int(totalItems),
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (int(totalItems) + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all books", responseData, meta))
	}
}
