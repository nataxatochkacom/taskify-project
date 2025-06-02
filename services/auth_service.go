package services

import (
	"errors"
	"task/TASKIFY-SAAS/models"
	"task/TASKIFY-SAAS/repositories"
)

// Регистрация нового пользователя
func Register(email, password string) error {
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := repositories.FindUserByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("пользователь с таким email уже существует")
	}

	// Хешируем пароль
	hashedPassword, err := models.HashPassword(password)
	if err != nil {
		return err
	}

	// Создаем нового пользователя
	newUser := &models.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}

	// Сохраняем пользователя в базу данных
	return repositories.CreateUser(newUser)
}

// Авторизация пользователя
func Login(email, password string) (*models.User, error) {
	// Находим пользователя по email
	user, err := repositories.FindUserByEmail(email)
	if err != nil || user == nil {
		return nil, errors.New("некорректные учетные данные")
	}

	// Проверяем пароль
	if !models.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("некорректные учетные данные")
	}

	return user, nil
}
