package main

import (
	"fmt"
	"task/TASKIFY-SAAS/handlers"
	"task/TASKIFY-SAAS/middlewares"
	"task/TASKIFY-SAAS/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // Глобальная переменная для доступа к базе данных

func initDB() {
	// Настройки подключения к базе данных
	dsn := "root:Dima1305@tcp(127.0.0.1:3306)/taskify?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Не удалось подключиться к базе данных!")
	}

	fmt.Println("Подключение к базе данных успешно!")

	// Инициализация репозитория пользователей
	repositories.InitUserRepository(db)
}

func main() {
	initDB() // Инициализация подключения к базе данных

	app := fiber.New()

	// Подключение статических файлов
	app.Static("/", "./static")

	// Роуты для регистрации и авторизации
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)

	// Защищенные роуты
	app.Get("/dashboard", middlewares.JwtProtected(), func(c *fiber.Ctx) error {
		return c.SendString("Добро пожаловать в личный кабинет!")
	})

	// Запуск сервера
	app.Listen(":3000")
}
