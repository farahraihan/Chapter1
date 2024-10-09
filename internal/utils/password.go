package utils

import "golang.org/x/crypto/bcrypt"

type PassUtilInterface interface {
	GeneratePassword(string) ([]byte, error)
	ComparePassword([]byte, []byte) error
}

type passUtil struct{}

func NewPassUtil() PassUtilInterface {
	return &passUtil{}
}

func (pu *passUtil) GeneratePassword(plainPass string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pu *passUtil) ComparePassword(currentPass, inputPass []byte) error {
	return bcrypt.CompareHashAndPassword(currentPass, inputPass)
}
