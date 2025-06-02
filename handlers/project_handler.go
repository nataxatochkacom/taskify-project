package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"task/TASKIFY-SAAS/models" 
)

type ProjectHandler struct {
	DB *gorm.DB
}

func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	var project models.Project

	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if result := h.DB.Create(&project); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create project",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}


func (h *ProjectHandler) GetProjects(c *fiber.Ctx) error {
	var projects []models.Project

	if err := h.DB.Find(&projects).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Ошибка при получении проектов",
		})
	}

	return c.JSON(projects)
}


func (h *ProjectHandler) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.DB.Delete(&models.Project{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Ошибка удаления проекта",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}


func (h *ProjectHandler) UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var input models.Project

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Неверный формат данных",
		})
	}

	var existing models.Project
	if err := h.DB.First(&existing, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Проект не найден",
		})
	}

	if err := h.DB.Model(&existing).Updates(input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Ошибка обновления проекта",
		})
	}

	return c.JSON(existing)
}
