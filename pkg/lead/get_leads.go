package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vadimegorov13/go-fiber-crm/pkg/common/models"
)

func (h handler) GetLeads(c *fiber.Ctx) error {
	var leads []models.Lead

	if result := h.DB.Find(&leads); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.JSON(&leads)
}
