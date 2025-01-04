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

func (uq *UserQuery) Login(email string) (users.User, error) {
	var result User
	err := uq.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.User{}, err
	}

	return result.ToUserEntity(), nil
}

func (uq *UserQuery) Register(newUsers users.User) error {
	cnvData := ToUserQuery(newUsers)
	err := uq.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

func (uq *UserQuery) UpdateUser(userID uint, updateUser users.User) error {
	cnvData := ToUserQuery(updateUser)

	qry := uq.db.Where("id = ?", userID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (uq *UserQuery) DeleteUser(userID uint) error {
	qry := uq.db.Where("id = ?", userID).Delete(&User{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (uq *UserQuery) GetUserByID(userID uint) (users.User, error) {
	var user users.User
	err := uq.db.First(&user, userID).Error
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func (uq *UserQuery) GetAllUsers(limit int, page int, search string) ([]users.User, int, error) {
	var usersList []User
	var totalItem int64

	offset := (page - 1) * limit

	err := uq.db.Model(&User{}).Where("name LIKE ?", "%"+search+"%").Count(&totalItem).Error
	if err != nil {
		return nil, 0, err
	}

	err = uq.db.Where("name LIKE ?", "%"+search+"%").Limit(limit).Offset(offset).Find(&usersList).Error
	if err != nil {
		return nil, 0, err
	}

	usersEntities := make([]users.User, len(usersList))
	for i, user := range usersList {
		usersEntities[i] = user.ToUserEntity()
	}

	return usersEntities, int(totalItem), nil
}

func (uq *UserQuery) IsAdmin(userID uint) (bool, error) {
	var user users.User
	if err := uq.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}

func (uq *UserQuery) AddPoints(userID uint, points uint) error {
	// Tambahkan poin ke pengguna berdasarkan userID
	result := uq.db.Model(&User{}).
		Where("id = ?", userID).
		Update("point", gorm.Expr("point + ?", points))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
