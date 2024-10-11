package handler

import (
	"chapter1/internal/features/users"
	"chapter1/internal/helpers"
	"chapter1/internal/utils"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UService
	tu  utils.JwtUtilityInterface
}

func NewUserHandler(s users.UService, t utils.JwtUtilityInterface) users.UHandler {
	return &UserHandler{
		srv: s,
		tu:  t,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		image, err := c.FormFile("image")
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

		var input RegisterRequest

		err = c.Bind(&input)
		if err != nil {
			log.Print("failed to bind register request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		newUser := RegisterToUser(input)

		err = uh.srv.Register(newUser, src, image.Filename)
		if err != nil {
			log.Print("failed to register user")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "registration unsuccessful", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "user registration successful", nil))
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("failed to bind login request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid login input", nil))
		}

		result, token, err := uh.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("login attempt failed")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "login unsuccessful", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "login successful", ToLoginResponse(result, token)))
	}
}

func (uh *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to update user")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		var src multipart.File
		var filename string
		image, err := c.FormFile("image")
		if err == nil {
			src, err = image.Open()
			if err != nil {
				log.Println("failed to open image file")
				return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "unable to process image", nil))
			}
			defer src.Close()
			filename = image.Filename
		}

		var req RegisterRequest
		err = c.Bind(&req)
		if err != nil {
			log.Println("failed to bind update request")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid input", nil))
		}

		updateUser := RegisterToUser(req)
		err = uh.srv.UpdateUser(uint(userID), updateUser, src, filename)
		if err != nil {
			log.Println("failed to update user")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "update unsuccessful", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user was successfully updated", nil))
	}
}

func (uh *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to delete user")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		memberIDStr := c.Param("id")
		memberID, err := strconv.ParseUint(memberIDStr, 10, 32)
		if err != nil {
			log.Println("invalid user ID")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid user ID", nil))
		}

		err = uh.srv.DeleteUser(uint(userID), uint(memberID))
		if err != nil {
			log.Println("failed to delete user")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "deletion unsuccessful", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user was successfully deleted", nil))
	}
}

func (uh *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to get user")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

		user, err := uh.srv.GetUserByID(uint(userID))
		if err != nil {
			log.Println("failed to get user by ID")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve user", nil))
		}

		UserResponse := ToUserResponse(user)

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user was successfully retrieved", UserResponse))

	}
}

func (uh *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {

		userID := uh.tu.DecodToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			log.Println("unauthorized attempt to get all users")
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFormat(http.StatusUnauthorized, "failed", "unauthorized", nil))
		}

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

		users, totalItems, err := uh.srv.GetAllUsers(uint(userID), limit, page, search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "failed to retrieve users", nil))
		}

		responseData := ToUserResponses(users)
		meta := helpers.Meta{
			TotalItems:   totalItems,
			ItemsPerPage: limit,
			CurrentPage:  page,
			TotalPages:   (totalItems + limit - 1) / limit,
		}

		return c.JSON(http.StatusOK, helpers.ResponseWithMetaFormat(http.StatusOK, "success", "successfully retrieved all users", responseData, meta))
	}
}
