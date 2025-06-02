package handlers

import (
	"github.com/gofiber/fiber/v2"
	"task/TASKIFY-SAAS/models"
	"gorm.io/gorm"
)
	
type TaskHandler struct {
    DB *gorm.DB
}


func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
    }

    if err := h.DB.Create(&task).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "DB error"})
    }

    return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
    var tasks []models.Task
    if err := h.DB.Find(&tasks).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "DB error"})
    }

    return c.JSON(tasks)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.DB.Delete(&models.Task{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Ошибка удаления задачи",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var input models.Task

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Неверные данные"})
	}

	if err := h.DB.Model(&models.Task{}).Where("id_Task = ?", id).Updates(input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Ошибка обновления"})
	}

	return c.JSON(input)
}
