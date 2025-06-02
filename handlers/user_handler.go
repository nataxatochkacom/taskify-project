package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"task/TASKIFY-SAAS/models"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	// Получаем JWT-токен из контекста Fiber
	userToken := c.Locals("user").(*jwt.Token)

	// Достаём claims из токена
	claims := userToken.Claims.(jwt.MapClaims)

	// Получаем userID из поля Issuer (строка → int)
	userIDStr, ok := claims["iss"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid user ID in token",
		})
	}

	// Загружаем пользователя из БД
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Возвращаем пользователя
	return c.JSON(user)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := claims["iss"].(string)

	var input struct {
		Full_name string `json:"full_name"`
		Email     string `json:"email"`
		Role      string `json:"role"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
	}

	var user models.User
	if err := h.DB.First(&user, "id_User = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	user.Full_name = input.Full_name
	user.Email = input.Email
	user.Role = input.Role

	if err := h.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "DB error"})
	}

	return c.JSON(user)
}

