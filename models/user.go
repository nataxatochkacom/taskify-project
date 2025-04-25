package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint   `gorm:"column:id_User;primaryKey" json:"id"`
	Full_name    string `gorm:"column:Full_name" json:"full_name"`
	Email        string `gorm:"column:Email;unique" json:"email"`
	PasswordHash string `gorm:"column:password_hash" json:"-"`
}

// Метод для хеширования пароля
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Метод для проверки пароля
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
