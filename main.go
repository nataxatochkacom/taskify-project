package main

import (
	"fmt"
	"task/TASKIFY-SAAS/handlers"
	"task/TASKIFY-SAAS/middlewares"
	"task/TASKIFY-SAAS/models"
	"task/TASKIFY-SAAS/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // Глобальная переменная для доступа к базе данных

func initDB() {
	dsn := "taskifyuser:taskifypass@tcp(mysql:3306)/taskify?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных!")
	}
	fmt.Println("Подключение к базе данных успешно!")

	// Миграция таблицы проектов
	if err := db.AutoMigrate(&models.Project{}); err != nil {
		panic("Ошибка миграции таблицы Project: " + err.Error())
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
    	panic("Ошибка миграции таблицы User: " + err.Error())
	}
	if err := db.AutoMigrate(&models.Task{}); err != nil {
    	panic("Ошибка миграции Task: " + err.Error())
	}
	if err := db.AutoMigrate(&models.SharedAccess{}); err != nil {
    	panic("Ошибка миграции SharedAccess: " + err.Error())
	}

	// Инициализация репозиториев
	repositories.InitUserRepository(db)
}

func main() {
	initDB()

	app := fiber.New()

	// Подключение статических файлов
	app.Static("/", "./static")

	// Роуты для регистрации и авторизации
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)

	// Роут для создания проекта (API)
	projectHandler := handlers.ProjectHandler{DB: db}
	app.Post("/api/projects", projectHandler.CreateProject)

	// Защищенные роуты
	userHandler := handlers.UserHandler{DB: db}
	app.Get("/api/user", middlewares.JwtProtected(), userHandler.GetCurrentUser)
	app.Get("/api/projects", projectHandler.GetProjects)
	app.Delete("/api/projects/:id", projectHandler.DeleteProject)
	app.Put("/api/projects/:id", projectHandler.UpdateProject)

	taskHandler := handlers.TaskHandler{DB: db}
	app.Post("/api/tasks", taskHandler.CreateTask)
	app.Get("/api/tasks", taskHandler.GetTasks)
	app.Delete("/api/tasks/:id", taskHandler.DeleteTask)
	app.Put("/api/tasks/:id", taskHandler.UpdateTask)


	responsibleHandler := handlers.ResponsibleHandler{DB: db}
	app.Get("/api/responsibles", responsibleHandler.GetResponsibles)

	userHandler = handlers.UserHandler{DB: db}
	app.Put("/api/user", middlewares.JwtProtected(), userHandler.UpdateProfile)

	shareHandler := handlers.ShareHandler{DB: db}
	app.Post("/api/share", middlewares.JwtProtected(), shareHandler.CreateShareLink)
	app.Get("/api/share/:uuid", shareHandler.GetSharedTasks)


	// Запуск сервера
	app.Listen(":8080")
}
