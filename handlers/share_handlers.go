package handlers

import (
    "net/http"
    "task/TASKIFY-SAAS/models"
    

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"
    "gorm.io/gorm"
    
)

type ShareHandler struct {
    DB *gorm.DB
}

type ShareRequest struct {
    Role string `json:"role"`
}

func (h *ShareHandler) CreateShareLink(c *fiber.Ctx) error {
    // 1. Достаём токен и user_id
    token := c.Locals("user").(*jwt.Token)
    claims := token.Claims.(jwt.MapClaims)

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Некорректный токен (user_id отсутствует)",
        })
    }
    userID := uint(userIDFloat)

    // 2. Парсим JSON-тело запроса
    var body struct {
        Role string `json:"role"`
    }

    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Некорректный формат запроса",
        })
    }

    // 3. Валидация роли
    validRoles := map[string]bool{
        "admin": true, "manager": true, "user": true, "viewer": true,
    }
    if !validRoles[body.Role] {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Недопустимая роль",
        })
    }

    // 4. Создаём UUID и сохраняем в БД
    shareUUID := uuid.New().String()

    share := models.SharedAccess{
        UUID:    shareUUID,
        Role:    body.Role,
        OwnerID: userID,
    }

    if err := h.DB.Create(&share).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Не удалось создать ссылку доступа",
        })
    }

    // 5. Возвращаем клиенту
    return c.JSON(fiber.Map{
        "link": "/share/" + shareUUID,
    })
}

func (h *ShareHandler) GetSharedTasks(c *fiber.Ctx) error {
    uuid := c.Params("uuid")
    var shared models.SharedAccess

    if err := h.DB.Where("uuid = ?", uuid).First(&shared).Error; err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Ссылка не найдена"})
    }

    var tasks []models.Task
    if err := h.DB.Where("id_User = ?", shared.OwnerID).Find(&tasks).Error; err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось получить задачи"})
    }

    return c.JSON(fiber.Map{
        "role":  shared.Role,
        "tasks": tasks,
    })
}
