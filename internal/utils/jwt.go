package utils

import (
	config "chapter1/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtUtilityInterface interface {
	GenereteJwt(id uint) (string, error)
	DecodToken(token *jwt.Token) float64
	DecodTokenV2(c echo.Context) (uint, error)
}

type JwtUtility struct{}

func NewJwtUtility() JwtUtilityInterface {
	return &JwtUtility{}
}

func (ju *JwtUtility) GenereteJwt(id uint) (string, error) {

	jwtKey := config.ImportSetting().JWTSecret
	data := jwt.MapClaims{}

	data["id"] = id
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Minute * 45).Unix()

	processToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	result, err := processToken.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return result, nil
}

func (ju *JwtUtility) DecodToken(token *jwt.Token) float64 {
	var result float64

	claim := token.Claims.(jwt.MapClaims)

	for _, val := range claim {
		fmt.Println(val)
	}

	if value, found := claim["id"]; found {
		result = value.(float64)
	}

	return result
}

func (ju *JwtUtility) DecodTokenV2(c echo.Context) (uint, error) {
	var result float64

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("error get jwt")
	}
	claim := token.Claims.(jwt.MapClaims)

	for _, val := range claim {
		fmt.Println(val)
	}

	if value, found := claim["id"]; found {
		result = value.(float64)
	} else {
		return 0, errors.New("error get jwt id")

	}

	return uint(result), nil
}
