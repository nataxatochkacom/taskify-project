package repositories

import (
	"task/TASKIFY-SAAS/models"
	"gorm.io/gorm"
)

var DB *gorm.DB // Глобальная переменная для доступа к базе данных

// Инициализация репозитория
func InitUserRepository(db *gorm.DB) {
	DB = db
}

// Создание нового пользователя
func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

// Поиск пользователя по email
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := DB.Where("Email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
