package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vadimegorov13/go-fiber-crm/pkg/common/db"
	"github.com/vadimegorov13/go-fiber-crm/pkg/lead"
)

func setUpLogger(app *fiber.App) {
	file, err := os.OpenFile("../logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Simple CRM using fiber",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	setUpLogger(app)

	dbHandler := db.InitDB()
	lead.RegisterRoutes(app, dbHandler)

	log.Fatal(app.Listen(":3000"))
}
