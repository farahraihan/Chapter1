package service

import (
	"chapter1/internal/features/users"
	"chapter1/internal/utils"
	"errors"
	"log"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	qry        users.UQuery
	pwd        utils.PassUtilInterface
	jwt        utils.JwtUtilityInterface
	cloudinary utils.CloudinaryUtilityInterface
}

func NewUserServices(q users.UQuery, p utils.PassUtilInterface, j utils.JwtUtilityInterface, c utils.CloudinaryUtilityInterface) users.UService {
	return &UserServices{
		qry:        q,
		pwd:        p,
		jwt:        j,
		cloudinary: c,
	}
}

func (us *UserServices) Login(email string, password string) (users.User, string, error) {

	result, err := us.qry.Login(email)

	if err != nil {
		log.Println("login sql error: ", err.Error())
		return users.User{}, "", errors.New("error in server")
	}

	err = us.pwd.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Println("Invalid password", err)
		return users.User{}, "", errors.New(bcrypt.ErrMismatchedHashAndPassword.Error())
	}

	token, err := us.jwt.GenereteJwt(result.ID)
	if err != nil {
		log.Println("Error On Jwt ", err)
		return users.User{}, "", errors.New("error on JWT")
	}

	return result, token, nil
}

func (us *UserServices) Register(newUsers users.User, src multipart.File, filename string) error {

	hashPw, err := us.pwd.GeneratePassword(newUsers.Password)
	if err != nil {
		log.Println("register generate password error", err.Error())
		return err
	}

	newUsers.Password = string(hashPw)
	newUsers.IsAdmin = false

	imageURL, err := us.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		return errors.New("failed to upload image")
	}
	newUsers.Image = imageURL

	err = us.qry.Register(newUsers)
	if err != nil {
		log.Println("register sql error:", err.Error())
		return errors.New("error in server")
	}

	return nil
}
