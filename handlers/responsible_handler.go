package handlers

import (
    "task/TASKIFY-SAAS/models"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type ResponsibleHandler struct {
    DB *gorm.DB
}

func (h *ResponsibleHandler) GetResponsibles(c *fiber.Ctx) error {
    var responsibles []models.Responsible
    if err := h.DB.Find(&responsibles).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Ошибка получения ответственных",
        })
    }
    return c.JSON(responsibles)
}
