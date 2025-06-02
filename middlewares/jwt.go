package middlewares

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v2" // Переименуем импорт для ясности
)

func JwtProtected() func(*fiber.Ctx) error {
    return jwtware.New(jwtware.Config{
        SigningKey: []byte("your-secret-key"), // Замени на свой секретный ключ
        ContextKey: "user", // <--- ОБЯЗАТЕЛЬНО!
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Unauthorized",
            })
        },
    })
}
