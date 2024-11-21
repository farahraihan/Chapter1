package service

import (
	"chapter1/internal/features/users"
	"chapter1/internal/utils"
	"errors"
	"log"
	"mime/multipart"
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
		log.Println("login query error: ", err.Error())
		return users.User{}, "", errors.New("login failed, please try again later")
	}

	err = us.pwd.ComparePassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Println("invalid password", err)
		return users.User{}, "", errors.New("invalid credentials")
	}

	token, err := us.jwt.GenerateJwt(result.ID)
	if err != nil {
		log.Println("error generating jwt", err)
		return users.User{}, "", errors.New("login failed, please try again later")
	}

	return result, token, nil
}

func (us *UserServices) Register(newUsers users.User, src multipart.File, filename string) error {

	hashPw, err := us.pwd.GeneratePassword(newUsers.Password)
	if err != nil {
		log.Println("register password generation error: ", err)
		return errors.New("registration failed, please try again later")
	}

	newUsers.Password = string(hashPw)
	newUsers.IsAdmin = false

	imageURL, err := us.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		log.Println("image upload failed: ", err)
		return errors.New("failed to upload image, please try again later")
	}
	newUsers.Image = imageURL

	err = us.qry.Register(newUsers)
	if err != nil {
		log.Println("register query error: ", err)
		return errors.New("registration failed, please try again later")
	}

	return nil
}

func (us *UserServices) UpdateUser(userID uint, updatedUser users.User, src multipart.File, filename string) error {

	if updatedUser.Password != "" {
		hashPassword, err := us.pwd.GeneratePassword(updatedUser.Password)
		if err != nil {
			log.Println("update password generation error: ", err)
			return errors.New("update failed, please try again later")
		}

		updatedUser.Password = string(hashPassword)
	}

	imageURL, err := us.cloudinary.UploadToCloudinary(src, filename)
	if err != nil {
		log.Println("image upload failed: ", err)
		return errors.New("failed to upload image, please try again later")
	}
	updatedUser.Image = imageURL

	err = us.qry.UpdateUser(userID, updatedUser)
	if err != nil {
		log.Println("update user query error: ", err)
		return errors.New("update failed, please try again later")
	}
	return nil
}

func (us *UserServices) DeleteUser(userID uint, memberID uint) error {
	isAdmin, err := us.qry.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("delete user permission error: ", err)
		return errors.New("access denied")
	}

	err = us.qry.DeleteUser(memberID)
	if err != nil {
		log.Println("delete user query error: ", err)
		return errors.New("delete failed, please try again later")
	}

	return nil
}

func (us *UserServices) GetUserByID(userID uint) (users.User, error) {
	user, err := us.qry.GetUserByID(userID)
	if err != nil {
		log.Println("get user by ID query error: ", err)
		return users.User{}, errors.New("failed to retrieve user, please try again later")
	}

	return user, nil
}

func (us *UserServices) GetAllUsers(userID uint, limit int, page int, search string) ([]users.User, int, error) {
	isAdmin, err := us.qry.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("get all users permission error: ", err)
		return nil, 0, errors.New("access denied")
	}

	users, totalItems, err := us.qry.GetAllUsers(limit, page, search)

	if err != nil {
		log.Println("get all users query error: ", err)
		return nil, 0, errors.New("failed to retrieve users, please try again later")
	}

	return users, totalItems, nil
}

func (us *UserServices) IsAdmin(userID uint) (bool, error) {
	return us.qry.IsAdmin(userID)
}
