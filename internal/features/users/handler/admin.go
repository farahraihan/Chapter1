package handler

import (
	"chapter1/internal/features/users"
	"chapter1/internal/helpers"
	"chapter1/internal/utils"
	"log"
	"net/http"

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
			log.Println("error get image file")
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "invalid image file", nil))
		}

		src, err := image.Open()
		if err != nil {
			log.Println("error open the image file")
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "failed", "Failed to open image file", nil))
		}
		defer src.Close()

		var input RegisterRequest

		err = c.Bind(&input)
		if err != nil {
			log.Print("error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "user register failed", nil))
		}

		// Convert request to reward model
		newUser := RegisterToUser(input, "")

		err = uh.srv.Register(newUser, src, image.Filename)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occurred", nil))
		}

		return c.JSON(http.StatusCreated, helpers.ResponseFormat(http.StatusCreated, "success", "user register successful", nil))
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest

		err := c.Bind(&input)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helpers.ResponseFormat(http.StatusBadRequest, "failed", "user login failed", nil))
		}

		result, token, err := uh.srv.Login(input.Email, input.Password)
		if err != nil {
			log.Print("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helpers.ResponseFormat(http.StatusInternalServerError, "error", "an unexpected error occured", nil))
		}

		return c.JSON(http.StatusOK, helpers.ResponseFormat(http.StatusOK, "success", "user login successful", ToLoginResponse(result, token)))
	}
}
