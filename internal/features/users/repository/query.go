package repository

import (
	"chapter1/internal/features/users"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(connect *gorm.DB) users.UQuery {
	return &UserQuery{
		db: connect,
	}
}

func (um *UserQuery) Login(email string) (users.User, error) {
	var result User
	err := um.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.User{}, err
	}

	return result.ToUserEntity(), nil
}

func (am *UserQuery) Register(newUsers users.User) error {
	cnvData := ToUserQuery(newUsers)
	err := am.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}
