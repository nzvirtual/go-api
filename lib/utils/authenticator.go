package utils

import (
	"github.com/nzvirtual/go-api/lib/database/models"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(email string, password string) (*models.User, error) {
	user := models.User{}
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}
