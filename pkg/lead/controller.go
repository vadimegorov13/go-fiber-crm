package lead

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	lead := app.Group("/lead")

	lead.Get("/", h.GetLeads)
	lead.Get("/:id", h.GetLead)
	lead.Post("/", h.NewLead)
	lead.Delete("/:id", h.DeleteLead)
}
