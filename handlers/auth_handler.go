package handlers

import (
    "net/http"
    "strconv"
    "time"

    "task/TASKIFY-SAAS/services"
    "github.com/dgrijalva/jwt-go" // Используем стандартную библиотеку для JWT
    "github.com/gofiber/fiber/v2"
)

// Обработчик регистрации
func RegisterHandler(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
    }

    email := data["email"]
    password := data["password"]

    if email == "" || password == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "Email and password are required",
        })
    }

    // Вызываем сервис для регистрации
    err := services.Register(email, password)
    if err != nil {
        return c.Status(http.StatusConflict).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    return c.Status(http.StatusCreated).JSON(fiber.Map{
        "message": "User registered successfully",
        "user": fiber.Map{
            "email": email,
        },
    })
}

// Обработчик входа
func LoginHandler(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
    }

    email := data["email"]
    password := data["password"]

    if email == "" || password == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "message": "Email and password are required",
        })
    }

    // Вызываем сервис для входа
    user, err := services.Login(email, password)
    if err != nil {
        return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    // Генерируем JWT токен
    token := generateJWT(user.ID, user.Email)

    return c.Status(http.StatusOK).JSON(fiber.Map{
        "message": "Login successful",
        "token":   token,
    })
}

// Функция для генерации JWT токена
func generateJWT(userID uint, email string) string {
    // Создаем claims (утверждения) для токена
    claims := jwt.StandardClaims{
        Issuer:    strconv.Itoa(int(userID)), // ID пользователя
        Subject:   email,                     // Email пользователя
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Время истечения токена
    }

    // Создаем новый токен с HMAC подписью
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Подписываем токен с использованием секретного ключа
    signedToken, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        // Обработка ошибки подписи токена
        return ""
    }

    return signedToken
}