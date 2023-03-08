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

	api := app.Group("/api")
	v1 := api.Group("/v1")
	lead := v1.Group("/lead")

	lead.Get("/", h.GetLeads)
	lead.Get("/:id", h.GetLead)
	lead.Post("/", h.NewLead)
	lead.Delete("/:id", h.DeleteLead)
}
