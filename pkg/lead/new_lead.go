package lead

import (
	"github.com/goccy/go-json"
	"github.com/vadimegorov13/go-fiber-crm/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) NewLead(c *fiber.Ctx) error {
	body := c.Body()

	var lead models.Lead

	if err := json.Unmarshal(body, &lead); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if result := h.DB.Create(&lead); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.JSON(&lead)
}
