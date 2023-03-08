package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vadimegorov13/go-fiber-crm/pkg/common/models"
)

func (h handler) DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	var lead models.Lead

	if result := h.DB.First(&lead, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&lead)

	return c.SendStatus(200)
}
