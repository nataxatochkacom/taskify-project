package middlewares

import (
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v2" // Переименуем импорт для ясности
)

func JwtProtected() func(*fiber.Ctx) error {
    return jwtware.New(jwtware.Config{ // Используем переименованный импорт
        SigningKey: []byte("your-secret-key"), // Замени на свой секретный ключ
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Unauthorized",
            })
        },
    })
}